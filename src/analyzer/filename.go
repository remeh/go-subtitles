// File analyzer to improve matches.
//
// Filename analyzer.
//
// Copyright © 2014 - Rémy MATHIEU

package analyzer

import (
	"log"
)

const (
	SERIE_SEASON_EPISODE_PATTERN = "S([0-9]{1,2})E([0-9]{1,2})"
)

// Analyzes the filename to return a analyze result.
func AnalyzeFilename(filename string) Result {
	filtered := GetFilename(filename)
	filtered = RemoveExtension(filtered)
	filtered = RemoveSpecialChars(filtered)
	season, episode := LookForSerieInfo(filtered)

	log.Println("Cleaned name:", filtered)

	isSerie := (season != 0 && episode != 0)

	return Result{
		Name:    filtered,
		IsSerie: isSerie,
		Season:  season,
		Episode: episode,
	}
}
