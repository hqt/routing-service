package server

import (
	"github.com/hqt/routing-service/src/routing"
	"log"
	"net/http"
	"net/http/httptest"
)

const (
	defaultServer = "0.0.0.0:3000"
)

// App represents  app struct
type App struct {
	mux   *http.ServeMux
	sever *http.Server
	graph routing.Graph
}

// NewApp returns App object
func NewApp(csvPath string) (*App, error) {
	stations, err := routing.ParseCSVToStations(csvPath)
	if err != nil {
		return nil, err
	}
	graph := routing.BuildGraph(stations)

	a := &App{
		mux:   http.NewServeMux(),
		graph: graph,
	}

	a.setUpRoute()
	return a, nil
}

func (app *App) setUpRoute() {
	app.mux.HandleFunc("/api/simple_route", app.simpleRoute)
	app.mux.HandleFunc("/api/advanced_route", app.advancedRoute)
}

// Start starts the application
func (app *App) Start() error {

	log.Println("start server on port 3000...")
	server := &http.Server{
		Addr:    defaultServer,
		Handler: app.mux,
	}
	app.sever = server
	return app.sever.ListenAndServe()
}

// Stop stops the application
func (app *App) Stop() error {
	return app.sever.Close()
}

// StartTestServer returns a new test server for testing
func (app *App) StartTestServer() (*httptest.Server, error) {
	return httptest.NewServer(app.mux), nil
}
