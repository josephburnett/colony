package colony

import "encoding/gob"

var _ Object = &Rock{}

func init() {
	gob.Register(&Rock{})
}

type Rock struct {
	Lifetime int
}

func NewRock() *Rock {
	return &Rock{Lifetime: 1000}
}

func (r *Rock) Owner() Owner {
	return Owner("")
}

func (r *Rock) Tick() {
	r.Lifetime = r.Lifetime - 1
}

func (r *Rock) Dead() bool {
	return r.Lifetime == 0
}

func (r *Rock) View(o Owner) *ObjectView {
	return &ObjectView{
		Type: "rock",
	}
}
