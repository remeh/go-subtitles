// Common helper for analyzing.
//
// Copyright © 2014 - Rémy MATHIEU

package analyzer

import (
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// Reads an absolute uri to a file to returns only the filename.
func GetFilename(filename string) string {
	return strings.ToUpper(filepath.Base(filename))
}

// Replaces common words or sysonym for a better matching.
func ReplaceCommonWords(filename string) string {
	synonymous := make(map[string]string)

	synonymous["BLU RAY"] = "BDRip"
	synonymous["BLURAY"] = "BDRip"

	for k, v := range synonymous {
		filename = strings.Replace(filename, k, v, -1)
	}

	return filename
}

// Looks for season / episode information
func LookForSerieInfo(filename string) (int, int) {
	r := regexp.MustCompile(SERIE_SEASON_EPISODE_PATTERN)
	matches := r.FindSubmatch([]byte(filename))
	if len(matches) == 3 {
		season, err := strconv.Atoi(string(matches[1]))
		if err != nil { // should never happen as the regexp has matched
			return 0, 0
		}
		episode, err := strconv.Atoi(string(matches[2]))
		if err != nil { // should never happen as the regexp has matched
			return 0, 0
		}
		return season, episode
	}
	return 0, 0
}

// Removes the extension from the file.
func RemoveExtension(filename string) string {
	// TODO Try to look if it's a directory, if so,
	// TODO do not remove any extension.
	for i := len(filename) - 1; i >= 0 && !os.IsPathSeparator(filename[i]); i-- {
		if filename[i] == '.' {
			return filename[:i]
		}
	}
	return filename
}

// Remove the special chars from the given string
// and returns the result.
func RemoveSpecialChars(filename string) string {
	// List of special characters
	specialChars := ".-_\"'()[]{}!%#"

	// Copy the string
	result := filename

	for i := 0; i < len(specialChars); i++ {
		result = strings.Replace(result, string(specialChars[i]), " ", -1)
	}

	return result
}
