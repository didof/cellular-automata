package main

type Cell interface {
	XY() (int, int)
	Alive() bool
	Set(state bool)
}

type SimpleCell struct {
	x, y  int
	alive bool
}

func (c *SimpleCell) XY() (int, int) {
	return c.x, c.y
}

func (c *SimpleCell) Alive() bool {
	return c.alive
}

func (c *SimpleCell) Set(state bool) {
	c.alive = state
}
