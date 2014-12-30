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

	Filename    string
	SubLanguage string
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

	log.Println("Searching for :", cliParams.Filename)
	err = client.Search(cliParams.Filename, cliParams.SubLanguage, 15)
	if err != nil {
		log.Println("Error while searching:", err)
	}

	client.LogOut()
	log.Println("Logged out.")
}

// Parse the CLI parameters.
func parseFlags() Parameters {
	username := flag.String("u", "", "Username on OpenSubtitles.org")
	password := flag.String("p", "", "Password on OpenSubtitles.org")
	language := flag.String("l", "en", "Language")
	subLanguage := flag.String("sl", "eng", "Subtitle language")
	useragent := flag.String("k", "OSTestUserAgent", "OpenSubtitles Registered User Agent")
	filename := flag.String("f", "", "Search for the given filename")

	flag.Parse()

	return Parameters{
		Username:    *username,
		Password:    *password,
		Language:    *language,
		SubLanguage: *subLanguage,
		UserAgent:   *useragent,
		Filename:    *filename,
	}
}
