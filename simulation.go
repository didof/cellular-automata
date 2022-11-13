package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Simulation interface {
	Init(populationPercentage float64)
	Process()
	Exit() <-chan struct{}
	Cells() []Cell
}

// Default implementation
type DefaultSimulation struct {
	width, height int
	Units         []Cell
}

func NewSimulation(width, height int) Simulation {
	return &DefaultSimulation{
		width:  width,
		height: height,
		Units:  make([]Cell, width*height),
	}
}

func (s *DefaultSimulation) Init(populationPercentage float64) {
	rand.Seed(time.Now().UnixNano())

	n := 0
	for x := 0; x < s.width; x++ {
		for y := 0; y < s.height; y++ {
			alive := rand.Float64() < populationPercentage
			s.Units[n] = &SimpleCell{x: uint(x), y: uint(y), alive: alive}
		}
	}
}

func (s *DefaultSimulation) Exit() <-chan struct{} {
	done := make(chan struct{})

	go func() {
		defer close(done)

		fmt.Println("cleanup")
	}()

	return done
}

func (s *DefaultSimulation) Process() {
}

func (s *DefaultSimulation) Cells() []Cell {
	return s.Units
}
