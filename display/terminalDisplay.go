package display

import (
	"fmt"
	"log"
)

type TerminalDisplay struct {
	width, height int
	i             int
}

func NewTerminalDisplay(width, height int) *TerminalDisplay {
	return &TerminalDisplay{width, height, 0}
}

const (
	ansiEscapeSeq = "\033c\x0c"
	brownSquare   = "\xF0\x9F\x9F\xAB"
	greenSquare   = "\xF0\x9F\x9F\xA9"
)

func (t TerminalDisplay) Clean() {
	fmt.Print(ansiEscapeSeq)
}

func (t *TerminalDisplay) Done() {
	t.i = 0
	fmt.Println("")
}

func (t *TerminalDisplay) DrawAlive() {
	if t.i > t.width*t.height {
		log.Fatal("out of display")
	} else if t.i%t.width == 0 {
		fmt.Println("")
	}

	fmt.Print(greenSquare)
	t.i++
}

func (t *TerminalDisplay) DrawDead() {
	if t.i > t.width*t.height {
		log.Fatal("out of display")
	} else if t.i%t.width == 0 {
		fmt.Println("")
	}

	fmt.Print(brownSquare)

	t.i++
}
