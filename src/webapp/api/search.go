// Search the best subtitle for the given filename.
// Returns an ordered (better first) list of found subtitles.
//
// Copyright © 2015 - Rémy MATHIEU

package api

import (
	"encoding/json"
	"net/http"

	"service"
	"webapp"
)

type Search struct {
	App *webapp.App
}

type SearchResponse struct {
	Subtitles []service.Subtitle `json:"subtitles"`
	Metadata  service.Metadata   `json:"metadata"`
}

func (h *Search) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO Log/save request

	// Parameters
	r.ParseForm()
	filename := r.Form.Get("f")
	if len(filename) == 0 {
		w.WriteHeader(400)
		return
	}
	language := r.Form.Get("l")
	if len(language) == 0 {
		language = "eng"
	}

	// Retrieve some subtitles
	subtitles, metadata, err := service.Search(h.App.Config, filename, language, 100)

	if err != nil {
		w.WriteHeader(400)
		return
	}

	response := SearchResponse{
		Subtitles: subtitles,
		Metadata:  metadata,
	}
	json, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
