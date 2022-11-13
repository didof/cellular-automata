package main

type Cell interface {
	XY() (int, int)
	Alive() bool
	Set(state bool)
}

type SimpleCell struct {
	x, y  uint
	alive bool
}

func (c *SimpleCell) XY() (int, int) {
	return int(c.x), int(c.y)
}

func (c *SimpleCell) Alive() bool {
	return c.alive
}

func (c *SimpleCell) Set(state bool) {
	c.alive = state
}

func NewCell(x, y uint, alive bool) SimpleCell {
	return SimpleCell{}
}
