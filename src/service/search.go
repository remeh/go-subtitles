package service

import (
	"log"
	"sort"

	"analyzer"
	"omdb"
	"opensubtitles"
	"webapp"
)

// TODO comment
// TODO language conversion
// TODO
func Search(config webapp.Config, filename string, language string, limit int) ([]Subtitle, Metadata, error) {
	subtitles := make([]Subtitle, 0)
	metadata := Metadata{}

	// Opens a client
	client := opensubtitles.NewOSClient(language, config.UserAgent)
	err := client.LogIn(config.Username, config.Password)

	if err != nil {
		log.Fatalf("Unable to connect to OpenSubtitles :\n%s\n", err)
		return subtitles, metadata, err
	}

	log.Println("Logged in, received token :", client.Token)
	// search
	log.Println("Searching for :", filename)
	response, err := client.Search(filename, language, limit)

	if err != nil {
		log.Println("Error while searching:", err)
		return subtitles, metadata, err
	}

	// logout
	client.LogOut()

	// Process the result
	for _, v := range response.SubtitleEntries {
		// convert to generic model
		sub := FromOSEntry(v)
		// compute score from filename
		sub.FilenameScore = analyzer.CompareFilenameSubtitleName(filename, sub.Filename)

		// append to the list but eliminate entries with -1 as filename score.
		if sub.FilenameScore > -1.0 {
			subtitles = append(subtitles, sub)
		}
	}

	// Metadata
	omdbClient := omdb.OMDBClient{}
	omdbResponse, err := omdbClient.Search(subtitles[0].IMDBId, "")
	if err == nil {
		metadata = FromOMDB(omdbResponse)
	}

	// Sort by Rating
	sort.Sort(Subtitles(subtitles))

	return subtitles, metadata, nil
}
