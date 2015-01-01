// Search the best subtitle for the given filename.
//
// Copyright © 2015 - Rémy MATHIEU

package api

import (
	"fmt"
	"net/http"

	"webapp"
)

type Search struct {
	App *webapp.App
}

type SearchResponse struct {
	Link string
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

	// TODO

	fmt.Fprintf(w, "TODO")
}
