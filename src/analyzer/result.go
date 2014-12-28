// File analyzer to improve matches.
//
// Filename analyzer.
//
// Copyright © 2014 - Rémy MATHIEU

package analyzer

type Result struct {
	Name    string // Cleared name of the file containing only the movie/serie name
	IsSerie bool   // True if this file looks like a serie file
	Season  int    // If it's a serie, which season it can be
	Episode int    // If it's an episode of a serie, which ep it would be.
}
