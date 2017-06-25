package colony

import (
	"math/rand"
)

type Owner string
type Point [2]int
type Direction [2]int
type Phermones map[Point]bool
type Objects map[Point]Object

func (p Point) Plus(d Direction) Point {
	return Point{p[0] + d[0], p[1] + d[1]}
}

func (p1 Point) Equals(p2 Point) bool {
	if p1[0] == p2[0] && p1[1] == p2[1] {
		return true
	}
	return false
}

type Object interface {
	Owner() Owner
	Point() Point
	Dead() bool
	View(Owner) *ObjectView
}

type AnimateObject interface {
	Move(Objects, Phermones) Point
	Attack(Object) bool
	TakeDamage(int)
	Strength() int
}

type World struct {
	owners    map[Owner]*Colony
	phermones map[Owner]Phermones
	objects   Objects
	colonies  map[Point]*Colony
	clients   Clients
}

func NewWorld() *World {
	return &World{
		owners:    make(map[Owner]*Colony),
		phermones: make(map[Owner]Phermones),
		objects:   make(Objects),
		colonies:  make(map[Point]*Colony),
	}
}

func (w *World) NewColony(o Owner) {
	var p Point
	for {
		p = Point{
			rand.Intn(40) - 20,
			rand.Intn(40) - 20,
		}
		if _, occupied := w.colonies[p]; !occupied {
			break
		}
	}
	c := &Colony{
		owner: o,
		point: p,
	}
	w.owners[o] = c
	w.phermones[o] = make(Phermones)
	w.colonies[p] = c
}

func (w *World) Advance() {
	objects := make([]Object, 0, len(w.objects))
	for _, o := range w.objects {
		objects = append(objects, o)
	}
	// Move objects
	perm := rand.Perm(len(objects))
	for _, i := range perm {
		o := objects[i]
		if o.Dead() {
			continue
		}
		if ao, ok := o.(AnimateObject); ok {
			fromPoint := o.Point()
			toPoint := ao.Move(w.objects, w.phermones[o.Owner()])
			if fromPoint.Equals(toPoint) {
				continue
			}
			target, occupied := w.objects[toPoint]
			if occupied {
				win := ao.Attack(target)
				if win {
					w.objects[toPoint] = o
				}
				delete(w.objects, fromPoint)
			} else {
				w.objects[toPoint] = o
				delete(w.objects, fromPoint)
			}
		}
	}
	// Produce objects
	for _, c := range w.colonies {
		ant, produced := c.Produce(w.objects)
		if produced {
			w.objects[ant.Point()] = ant
		}
	}
	// Remove the dead
	for _, o := range w.objects {
		if o.Dead() {
			delete(w.objects, o.Point())
		}
	}
}
