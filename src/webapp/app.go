// Application for the webapp.
//
// Copyright © 2015 - Rémy MATHIEU

package webapp

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Main webapp structure.
type App struct {
	Config Config
	Router *mux.Router
}

// Creates a new webapp and instanciate the router.
func NewApp(config Config) *App {
	return &App{
		Config: config,
		Router: mux.NewRouter(),
	}
}

// Starts the webapp.
func (a *App) Start(closeChannel chan int) {
	// Add the ultimate route to serve static files.
	a.Router.PathPrefix("/").Handler(http.FileServer(http.Dir(a.Config.StaticDirectory)))

	go http.ListenAndServe(a.Config.Addr, a.Router)
	<-closeChannel // Wait for something to close the app.
}
