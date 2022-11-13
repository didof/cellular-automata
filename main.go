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
var port *uint
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
	flag.Parse()
}

func usage() {
	fmt.Println("sim")
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	width := flag.Int("width", 3, "The world width")
	height := flag.Int("height", 3, "The world height")

	sim := NewSimulation(*width, *height)
	server := NewBrowserSimulationServer("0.0.0.0", *port, *public, *headless)
	server.AddTerminationSignals(os.Interrupt, syscall.SIGTERM)
	server.Serve(sim)
}

func Run(sim Simulation) {
	sim.Init()
}
