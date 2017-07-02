package colony

type Colony struct {
	owner   Owner
	point   Point
	produce bool
}

func (c *Colony) Owner() Owner {
	return c.owner
}

func (c *Colony) Point() Point {
	return c.point
}

func (c *Colony) Produce(o map[Point]Object) (*Ant, bool) {
	if c.produce {
		_, obstructed := o[c.Point()]
		if !obstructed {
			c.produce = false
			return &Ant{
				owner:     c.owner,
				point:     c.point,
				direction: RandomDirection(D_AROUND),
				speed:     5,
				strength:  2,
				endurance: 10,
			}, true
		}
	}
	return nil, false
}
