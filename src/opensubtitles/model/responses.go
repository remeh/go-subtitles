// Client for OpenSubtitles which use
// XML-RPC calls (sic.)
//
// Declaration of every used responses.
//
// Copyright © 2014 - Rémy MATHIEU

package model

import "fmt"

// The log in response with the auth token.
type LogInResponse struct {
	Token   string  `xmlrpc:"token"`
	Status  string  `xmlrpc:"status"`
	Seconds float32 `xmlrpc:"seconds"`
}

// The logout response.
type LogOutResponse struct {
	Status  string  `xmlrpc:"status"`
	Seconds float32 `xmlrpc:"seconds"`
}

// Response received when searching for subtitles.
type SearchSubtitlesResponse struct {
	Status          string          `xmlrpc:"status"`
	Seconds         float32         `xmlrpc:"seconds"`
	SubtitleEntries []SubtitleEntry `xmlrpc:"data"`
}

// One possible match of subtitle
type SubtitleEntry struct {
	MovieName        string
	MovieReleaseName string
	MovieKind        string // episode, movie, ...
	LanguageName     string
	SubRating        string
	SubFileName      string // To compare with the video filename
	SeriesSeason     string
	SeriesEpisode    string
	SubDownloadLink  string
	ZipDownloadLink  string
	SubtitlesLink    string
}

func (e SubtitleEntry) String() string {
	return fmt.Sprintf("SubtitleEntry: MovieName[%s] MovieReleaseName[%s] MovieKind[%s] LanguageName[%s] SubFileName[%s] SubRating[%s] SeriesSeason[%s] SeriesEpisode[%s] SubDownloadLink[%s] ZipDownloadLink[%s] SubtitlesLink[%s]\n", e.MovieName, e.MovieReleaseName, e.MovieKind, e.LanguageName, e.SubFileName, e.SubRating, e.SeriesSeason, e.SeriesEpisode, e.SubDownloadLink, e.ZipDownloadLink, e.SubtitlesLink)
}
