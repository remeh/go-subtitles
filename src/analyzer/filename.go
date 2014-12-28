// File analyzer to improve matches.
//
// Filename analyzer.
//
// Copyright © 2014 - Rémy MATHIEU

package analyzer

// Analyzes the filename to return a analyze result.
func AnalyzeFilename(filename string) Result {
	return Result{
		Name: filename,
	}
}
