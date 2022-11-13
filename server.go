package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"runtime"
	"strconv"
	"syscall"
)

type SimulationServer interface {
	Serve(sim Simulation)
	AddTerminationSignals(signals ...interface{})
}

type BrowserSimulationServer struct {
	sim               Simulation
	server            *http.Server
	terminationSignal chan os.Signal
	headless          bool
}

func (s *BrowserSimulationServer) Serve() error {
	fmt.Printf("Started simulation server at http://%s\n", s.server.Addr)

	if !s.headless {
		go open(fmt.Sprintf("http://%s", s.server.Addr))
	}

	return s.server.ListenAndServe()
}

func (s *BrowserSimulationServer) AddTerminationSignals(signals ...interface{}) {
	signal.Notify(s.terminationSignal, os.Interrupt, syscall.SIGTERM)
	go func() {
		defer close(s.terminationSignal)
		<-s.terminationSignal
		fmt.Print("\n")
		<-s.sim.Exit()
		fmt.Println("Shutting down")
		os.Exit(1)
	}()
}

func (s *BrowserSimulationServer) Handle(w http.ResponseWriter, r *http.Request) {
	s.sim.Process()

	width, height := s.sim.Sizes()
	grid := make([][]bool, height)

	c := 0
	cells := s.sim.Cells()

	for i := 0; i < height; i++ {
		grid[i] = make([]bool, width)
		for j := 0; j < width; j++ {
			grid[i][j] = cells[c].Alive()
			c++
		}
	}

	res, err := json.Marshal(grid)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(res)
}

func NewBrowserSimulationServer(sim Simulation, host string, port uint, publicPath, HTMLfile string, headless bool) BrowserSimulationServer {
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    host + ":" + strconv.Itoa(int(port)),
		Handler: mux,
	}

	mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir(publicPath))))

	templater := Templater{
		publicPath: publicPath,
		HTMLfile:   HTMLfile,
	}
	mux.HandleFunc("/", templater.Handle)

	s := BrowserSimulationServer{
		sim:               sim,
		server:            server,
		terminationSignal: make(chan os.Signal),
		headless:          headless,
	}
	mux.HandleFunc("/frame", s.Handle)

	return s
}

type Templater struct {
	publicPath string
	HTMLfile   string
}

func (t *Templater) Handle(w http.ResponseWriter, r *http.Request) {
	if tmpl, err := template.ParseFiles(filepath.Join(t.publicPath, t.HTMLfile)); err != nil {
		log.Fatal(err)
	} else {
		tmpl.Execute(w, nil)
	}
}

func open(url string) {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}
