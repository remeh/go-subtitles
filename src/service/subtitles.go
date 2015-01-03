// Containers.
//
// Copyright © 2015 - Rémy MATHIEU

package service

import (
	"fmt"

	"opensubtitles/model"
)

// Definition of a subtitle.
type Subtitle struct {
	MovieName        string
	MovieReleaseName string
	MovieKind        string // episode, movie, ...
	Language         string
	Rating           string
	Filename         string // To compare with the video filename
	SeriesSeason     string
	SeriesEpisode    string
	DownloadLink     string
	ZipDownloadLink  string
	FilenameScore    float32 // score computed from the filename
}

func (s Subtitle) String() string {
	return fmt.Sprintf("Subtitle: MovieName[%s] MovieReleaseName[%s] MovieKind[%s] Language[%s] Filename[%s] Rating[%s] SeriesSeason[%s] SeriesEpisode[%s] DownloadLink[%s] ZipDownloadLink[%s] FilenameScore[%f]\n", s.MovieName, s.MovieReleaseName, s.MovieKind, s.Language, s.Filename, s.Rating, s.SeriesSeason, s.SeriesEpisode, s.DownloadLink, s.ZipDownloadLink, s.FilenameScore)
}

// Converts the OS subtitle entry to a common Subtitle.
func FromOSEntry(e model.SubtitleEntry) Subtitle {
	return Subtitle{
		MovieName:        e.MovieName,
		MovieReleaseName: e.MovieReleaseName,
		MovieKind:        e.MovieKind,
		Language:         e.LanguageName,
		Rating:           e.SubRating,
		Filename:         e.SubFileName,
		SeriesSeason:     e.SeriesSeason,
		SeriesEpisode:    e.SeriesEpisode,
		DownloadLink:     e.SubDownloadLink,
		ZipDownloadLink:  e.ZipDownloadLink,
	}
}

// Implementing the sorting interface.
type Subtitles []Subtitle

func (s Subtitles) Len() int {
	return len(s)
}

// we invert the less for desc sort
func (s Subtitles) Less(i int, j int) bool {
	return s[i].FilenameScore > s[j].FilenameScore
}

func (s Subtitles) Swap(i int, j int) {
	s[i], s[j] = s[j], s[i]
}
