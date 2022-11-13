package main

import "fmt"

type Simulation interface {
	Init()
	Process()
	Exit() <-chan struct{}
}

// Default implementation
type Sim struct {
	width, height int
}

func NewSimulation(width, height int) Simulation {
	return &Sim{
		width:  width,
		height: height,
	}
}

func (s *Sim) Init() {
}

func (s *Sim) Exit() <-chan struct{} {
	done := make(chan struct{})

	go func() {
		defer close(done)

		fmt.Println("cleanup")
	}()

	return done
}

func (s *Sim) Process() {
}
