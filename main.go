package main

import (
	"flag"
	"log"

	"opensubtitles"
)

// CLI parameters.
type Parameters struct {
	Username  string
	Password  string
	Language  string
	UserAgent string
}

func main() {
	cliParams := parseFlags()

	// Opens a client
	client := opensubtitles.NewOSClient(cliParams.Language, cliParams.UserAgent)
	err := client.LogIn(cliParams.Username, cliParams.Password)

	if err != nil {
		log.Fatalf("Unable to connect to OpenSubtitles :\n%s\n", err)
	}

	log.Println("Logged in, received token :", client.Token)

	client.Search("Query search here", 1)

	/*
		if client.LogOut() != nil {
			log.Println("Error while logging out.")
		}
		log.Println("Logged out.")
	*/
}

func parseFlags() Parameters {
	username := flag.String("u", "", "Username on OpenSubtitles.org")
	password := flag.String("p", "", "Password on OpenSubtitles.org")
	language := flag.String("l", "en", "Language")
	useragent := flag.String("k", "OSTestUserAgent", "OpenSubtitles Registered User Agent")

	return Parameters{
		Username:  *username,
		Password:  *password,
		Language:  *language,
		UserAgent: *useragent,
	}
}
