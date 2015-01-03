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
	app.Router.Handle("/api/1.0/search", &api.Search{App: app})

	closeChannel := make(chan int)

	// Starts listening.
	app.Start(closeChannel)
}

func readFlags() webapp.Config {
	// Search config
	username := flag.String("u", "", "Username on OpenSubtitles.org")
	password := flag.String("p", "", "Password on OpenSubtitles.org")
	language := flag.String("l", "en", "Language")
	useragent := flag.String("k", "OSTestUserAgent", "OpenSubtitles Registered User Agent")

	// Webapp config
	addr := flag.String("addr", ":9000", "Address to listen to.")
	static := flag.String("dir", "static/", "The directory containing the static files to serve.")

	flag.Parse()

	return webapp.Config{
		Addr:            *addr,
		StaticDirectory: *static,

		Username:  *username,
		Password:  *password,
		Language:  *language,
		UserAgent: *useragent,
	}
}
