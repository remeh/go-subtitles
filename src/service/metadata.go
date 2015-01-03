// Containers.
//
// Copyright © 2015 - Rémy MATHIEU

package service

import (
	"omdb"
)

type Metadata struct {
	Title      string `json:"title"`
	Year       string `json:"year"`
	Type       string `json:"type"`
	Rated      string `json:"rated"`
	Released   string `json:"released"`
	Runtime    string `json:"runtime"`
	Genre      string `json:"genre"`
	Director   string `json:"director"`
	Writer     string `json:"writer"`
	Actors     string `json:"actors"`
	Plot       string `json:"plot"`
	Language   string `json:"language"`
	Country    string `json:"country"`
	Image      string `json:"image"`
	IMDBRating string `json:"imdbRating"`
	IMDBId     string `json:"imdbID`
}

func FromOMDB(response omdb.OMDBResponse) Metadata {
	return Metadata{
		Title:      response.Title,
		Year:       response.Year,
		Type:       response.Type,
		Rated:      response.Rated,
		Released:   response.Released,
		Runtime:    response.Runtime,
		Genre:      response.Genre,
		Director:   response.Director,
		Writer:     response.Writer,
		Actors:     response.Actors,
		Plot:       response.Plot,
		Language:   response.Language,
		Country:    response.Country,
		Image:      response.Poster,
		IMDBRating: response.IMDBRating,
		IMDBId:     response.IMDBId,
	}
}
