package service

import (
	"log"

	"opensubtitles"
	"opensubtitles/model"
)

// TODO comment
// TODO return a more abstract object ?
// TODO
func Search(filename string, limit int) []model.SubtitleEntry {
	// Opens a client
	client := opensubtitles.NewOSClient("eng", "OSTestUserAgent")
	err := client.LogIn("", "")

	if err != nil {
		log.Fatalf("Unable to connect to OpenSubtitles :\n%s\n", err)
	}

	log.Println("Logged in, received token :", client.Token)

	log.Println("Searching for :", filename)
	response, err := client.Search(filename, "eng", limit)

	if err != nil {
		log.Println("Error while searching:", err)
	}

	client.LogOut()

	return response.SubtitleEntries
}
