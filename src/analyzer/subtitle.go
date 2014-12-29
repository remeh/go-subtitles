// Subtitle analyzer.
//
// Copyright © 2014 - Rémy MATHIEU

package analyzer

import ()

// Analyzes the filename to return a analyze result.
func CompareFilenameSubtitleName(fileURI string, subtitleURI string) {
	filename := GetFilename(fileURI)
	sName := GetFilename(subtitleURI)

	// clears both filename
	filename = RemoveExtension(filename)
	filename = RemoveSpecialChars(filename)

	sName = RemoveExtension(sName)
	sName = RemoveSpecialChars(sName)

	// We now have two cleared file names.
	// The idea is to split by spaces and look how much of the "words"
	// of the subtitles filename we can find in the video filename

	// TODO Should be applied on MovieReleaseName and SubFileName
}
