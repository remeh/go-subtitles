// Stub for the go-subtitles API.
//
// Copyright © 2015 - Rémy MATHIEU

package api

import (
	"fmt"
	"net/http"

	"webapp"
)

type Index struct {
	App *webapp.App
}

func (h *Index) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}
