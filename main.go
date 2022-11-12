package main

import (
	"flag"

	"github.com/didof/conway-game-of-life/display"
	"github.com/didof/conway-game-of-life/game"
)

func main() {
	width := flag.Int("width", 10, "The world width")
	height := flag.Int("height", 10, "The world height")

	display := display.NewTerminalDisplay(*width, *height)
	world := game.NewWorld(*width, *height)
	world.Randomize(false)

	world.Display(display)
}
