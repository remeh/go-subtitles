// File analyzer to improve matches.
//
// Filename analyzer.
//
// Copyright © 2014 - Rémy MATHIEU

package analyzer

import (
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

const (
	SERIE_SEASON_EPISODE_PATTERN = "S([0-9]{1,2})E([0-9]{1,2})"
)

// Analyzes the filename to return a analyze result.
func AnalyzeFilename(filename string) Result {
	filtered := filepath.Base(filename)
	filtered = removeExtension(filtered)
	filtered = removeSpecialChars(filtered)
	season, episode := lookForSerieInfo(filtered)

	log.Println("Cleaned name:", filtered)

	isSerie := (season != 0 && episode != 0)

	return Result{
		Name:    filtered,
		IsSerie: isSerie,
		Season:  season,
		Episode: episode,
	}
}

// Looks for season / episode information
func lookForSerieInfo(filename string) (int, int) {
	r := regexp.MustCompile(SERIE_SEASON_EPISODE_PATTERN)
	matches := r.FindSubmatch([]byte(filename))
	if len(matches) == 3 {
		season, err := strconv.Atoi(string(matches[1]))
		if err != nil { // should never happen as the regexp has matched
			return 0, 0
		}
		episode, err := strconv.Atoi(string(matches[2]))
		if err != nil { // should never happen as the regexp has matched
			return 0, 0
		}
		return season, episode
	}
	return 0, 0
}

// Removes the extension from the file.
func removeExtension(filename string) string {
	for i := len(filename) - 1; i >= 0 && !os.IsPathSeparator(filename[i]); i-- {
		if filename[i] == '.' {
			return filename[:i]
		}
	}
	return filename
}

// Remove the special chars from the given string
// and returns the result.
func removeSpecialChars(filename string) string {
	// List of special characters
	specialChars := ".-_\"'()[]{}!%#"

	// Copy the string
	result := filename

	for i := 0; i < len(specialChars); i++ {
		result = strings.Replace(result, string(specialChars[i]), " ", -1)
	}

	return result
}
