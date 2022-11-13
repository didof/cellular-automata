package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
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
	terminationSignal chan os.Signal
	headless          bool
	server            *http.Server
}

func (s *BrowserSimulationServer) Serve(sim Simulation) error {
	s.sim = sim

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

func NewBrowserSimulationServer(host string, port uint, publicPath string, headless bool) BrowserSimulationServer {
	mux := http.NewServeMux()
	mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir(publicPath))))
	mux.HandleFunc("/", handleIndex)

	server := &http.Server{
		Addr:    host + ":" + strconv.Itoa(int(port)),
		Handler: mux,
	}

	return BrowserSimulationServer{
		sim:               nil,
		terminationSignal: make(chan os.Signal),
		headless:          headless,
		server:            server,
	}
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(*public + "/index.html")
	if err != nil {
		log.Fatal(err)
	}

	t.Execute(w, nil)
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
