package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/didof/conway-game-of-life/display"
	"github.com/didof/conway-game-of-life/game"
)

func main() {
	width := flag.Int("width", 10, "The world width")
	height := flag.Int("height", 10, "The world height")

	ctx, _ := context.WithCancel(context.Background())

	display := display.NewTerminalDisplay(*width, *height)
	world := game.NewWorld(*width, *height)
	world.Randomize(false)

	tick := time.NewTicker(time.Second)

	i := 0
gameloop:
	for {
		select {
		case <-ctx.Done():
			break gameloop
		case <-tick.C:
			world.Display(display)
			i++
			fmt.Printf("gen #%d", i)
		}
	}
}
