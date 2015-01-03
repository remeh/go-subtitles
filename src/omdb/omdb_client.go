// OMDBapi.com client to request
// for information on movies / series.
//
// Copyright © 2014 - Rémy MATHIEU

package omdb

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	OMDBAPI_URL = "http://www.omdbapi.com/"
)

type OMDBClient struct {
	httpClient http.Client
}

type OMDBResponse struct {
	Title      string
	Year       string
	Type       string
	Rated      string
	Released   string
	Runtime    string
	Genre      string
	Director   string
	Writer     string
	Actors     string
	Plot       string
	Language   string
	Country    string
	Poster     string
	Response   string
	Error      string
	IMDBRating string `json:"imdbRating"`
	IMDBId     string `json:"imdbID`
}

func (c OMDBClient) Search(imdbId string, movieType string) (OMDBResponse, error) {
	omdbResponse := OMDBResponse{}

	// Creates the request
	url := OMDBAPI_URL + "?i=tt" + imdbId

	// TODO adds movie type

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return omdbResponse, err
	}

	// HTTP call
	resp, err := c.httpClient.Do(req)

	if err != nil {
		return omdbResponse, err
	}

	defer resp.Body.Close()

	// Reads the body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return omdbResponse, err
	}

	log.Println(string(body))

	json.Unmarshal(body, &omdbResponse)

	return omdbResponse, nil
}
