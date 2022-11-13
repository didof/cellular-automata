package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"syscall"
)

var public, HTMLfile *string
var width, height, port *uint
var headless *bool
var ipp *float64

func init() {
	flag.Usage = usage
	d, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	public = flag.String("public", d+"/public", "path to public directory")
	HTMLfile = flag.String("HTMLfile", "automatic.html", "HTML file name (with extention)")
	port = flag.Uint("port", 7272, "the port where the simulation server starts")
	headless = flag.Bool("headless", false, "whether to run an headless simulation")
	width = flag.Uint("width", 10, "the width of the simulation world")
	height = flag.Uint("height", 10, "the height of the simulation world")
	ipp = flag.Float64("ipp", 0.3, "initial population percentage")
	flag.Parse()
}

func usage() {
	fmt.Println("sim")
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	sim := NewSimulation(int(*width), int(*height))
	sim.Init(*ipp)

	// sim.Process()

	server := NewBrowserSimulationServer(sim, "0.0.0.0", *port, *public, *HTMLfile, *headless)
	server.AddTerminationSignals(os.Interrupt, syscall.SIGTERM)
	server.Serve()
}
