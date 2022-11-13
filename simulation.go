package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Simulation interface {
	Init(populationPercentage float64)
	Process()
	Exit() <-chan struct{}

	Cells() []Cell
	Sizes() (int, int)
}

type DefaultSimulation struct {
	width, height    int
	Units            []Cell
	neighboursFinder *NeighboursFinder
}

func NewSimulation(width, height int) Simulation {
	return &DefaultSimulation{
		width:            width,
		height:           height,
		Units:            make([]Cell, width*height),
		neighboursFinder: NewNeighboursFinder(width, height),
	}
}

func (s *DefaultSimulation) Init(populationPercentage float64) {
	rand.Seed(time.Now().UnixNano())

	n := 0
	for x := 0; x < s.width; x++ {
		for y := 0; y < s.height; y++ {
			alive := rand.Float64() < populationPercentage
			s.Units[n] = &SimpleCell{x: x, y: y, alive: alive}
			n++
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
	var wg sync.WaitGroup

	wg.Add(len(s.Units))

	for index, unit := range s.Units {
		go func(index int, unit Cell) {
			defer wg.Done()
			neighbours := s.neighboursFinder.Find(index)
			var ncount int

			for _, neighbour := range neighbours {
				if s.Units[neighbour].Alive() {
					ncount++
				}
			}

			if unit.Alive() {
				if ncount < 2 || ncount > 3 {
					unit.Set(false)
				}
			} else {
				if ncount == 3 {
					unit.Set(true)
				}
			}
		}(index, unit)
	}

	wg.Wait()
}

func (s *DefaultSimulation) Cells() []Cell {
	return s.Units
}

func (s *DefaultSimulation) Sizes() (width int, height int) {
	return s.width, s.height
}
