// Containers.
//
// Copyright © 2015 - Rémy MATHIEU

package service

import (
	"fmt"
	"strconv"

	"opensubtitles/model"
)

// Definition of a subtitle.
type Subtitle struct {
	MovieName        string  `json:"movie_name"`
	MovieReleaseName string  `json:"movie_release_name"`
	MovieKind        string  `json:"movie_kind"` // episode, movie, ...
	Language         string  `json:"language"`
	Rating           string  `json:"rating"`
	Filename         string  `json:"filename"` // To compare with the video filename
	SeriesSeason     string  `json:"series_season"`
	SeriesEpisode    string  `json:"series_episode"`
	DownloadCount    int     `json:"download_count"`
	DownloadLink     string  `json:"download_link"`
	ZipDownloadLink  string  `json:"zip_download_link"`
	FilenameScore    float32 `json:"filename_score"` // score computed from the filename
}

func (s Subtitle) String() string {
	return fmt.Sprintf("Subtitle: MovieName[%s] MovieReleaseName[%s] MovieKind[%s] Language[%s] Filename[%s] Rating[%s] SeriesSeason[%s] SeriesEpisode[%s] DownloadLink[%s] ZipDownloadLink[%s] FilenameScore[%f]\n", s.MovieName, s.MovieReleaseName, s.MovieKind, s.Language, s.Filename, s.Rating, s.SeriesSeason, s.SeriesEpisode, s.DownloadLink, s.ZipDownloadLink, s.FilenameScore)
}

// Converts the OS subtitle entry to a common Subtitle.
func FromOSEntry(e model.SubtitleEntry) Subtitle {
	downloadCount, err := strconv.Atoi(e.SubDownloadsCnt)
	if err != nil {
		downloadCount = 0
	}
	return Subtitle{
		MovieName:        e.MovieName,
		MovieReleaseName: e.MovieReleaseName,
		MovieKind:        e.MovieKind,
		Language:         e.LanguageName,
		Rating:           e.SubRating,
		Filename:         e.SubFileName,
		SeriesSeason:     e.SeriesSeason,
		SeriesEpisode:    e.SeriesEpisode,
		DownloadCount:    downloadCount,
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
