package service

import (
	"log"
	"sort"

	"analyzer"
	"opensubtitles"
)

// TODO comment
// TODO language conversion
// TODO
func Search(filename string, language string, limit int) ([]Subtitle, error) {
	subtitles := make([]Subtitle, 0)

	// Opens a client
	client := opensubtitles.NewOSClient("eng", "OSTestUserAgent")
	err := client.LogIn("", "")

	if err != nil {
		log.Fatalf("Unable to connect to OpenSubtitles :\n%s\n", err)
		return subtitles, err
	}

	log.Println("Logged in, received token :", client.Token)
	// search
	log.Println("Searching for :", filename)
	response, err := client.Search(filename, language, limit)

	if err != nil {
		log.Println("Error while searching:", err)
		return subtitles, err
	}

	// logout
	client.LogOut()

	// Process the result
	for _, v := range response.SubtitleEntries {
		// convert to generic model
		sub := FromOSEntry(v)
		// compute score from filename
		sub.FilenameScore = analyzer.CompareFilenameSubtitleName(filename, sub.Filename)
		// append to the list
		subtitles = append(subtitles, sub)
	}

	// Sort by Rating
	sort.Sort(Subtitles(subtitles))

	return subtitles, nil
}
