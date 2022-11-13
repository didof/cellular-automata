package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"syscall"
)

var public *string
var width, height, port *uint
var headless *bool

func init() {
	flag.Usage = usage
	d, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	public = flag.String("public", d+"/public", "path to public directory")
	port = flag.Uint("port", 7272, "the port where the simulation server starts")
	headless = flag.Bool("headless", false, "whether to run an headless simulation")
	width = flag.Uint("width", 10, "the width of the simulation world")
	height = flag.Uint("height", 10, "the height of the simulation world")
	flag.Parse()
}

func usage() {
	fmt.Println("sim")
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	sim := NewSimulation(int(*width), int(*height))

	server := NewBrowserSimulationServer("0.0.0.0", *port, *public, *headless)
	server.AddTerminationSignals(os.Interrupt, syscall.SIGTERM)
	server.Serve(sim)
}

func Run(sim Simulation) {
	sim.Init()
}
