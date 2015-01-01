// Webapp of go-subtitles.
//
// Copyright © 2015 - Rémy MATHIEU

package main

import (
	"flag"

	"webapp"
	"webapp/api"
)

// Main for the webapp.
func main() {
	config := readFlags()
	app := webapp.NewApp(config)

	app.Router.Handle("/api/1.0/", &api.Index{App: app})

	closeChannel := make(chan int)

	// Starts listening.
	app.Start(closeChannel)
}

func readFlags() webapp.Config {
	var addr = flag.String("addr", ":9000", "Address to listen to.")
	var static = flag.String("dir", "static/", "The directory containing the static files to serve.")
	flag.Parse()

	return webapp.Config{
		Addr:            *addr,
		StaticDirectory: *static,
	}
}
