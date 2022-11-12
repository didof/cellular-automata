package game

import (
	"math/rand"
	"time"

	"github.com/didof/conway-game-of-life/display"
)

type World [][]Cell

func NewWorld(width, height int) World {
	w := make(World, height)
	for i := range w {
		w[i] = make([]Cell, width)
	}
	return w
}

func (w World) Randomize(resetSeed bool) {
	if resetSeed {
		rand.Seed(time.Now().UnixNano())
	}

	for _, row := range w {
		for i := range row {
			row[i].Alive = rand.Intn(4) == 1
		}
	}
}

func (w World) Display(d display.Display) {
	d.Clean()
	defer d.Done()
	for _, row := range w {
		for _, cell := range row {
			switch {
			case cell.Alive:
				d.DrawAlive()
			default:
				d.DrawDead()
			}
		}
	}
}
