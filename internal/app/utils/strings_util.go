package utils

import (
	"regexp"
)

func StripString(s string) string {
	re := regexp.MustCompile(`\\n|\\`)
	cleanedString := re.ReplaceAllString(s, "")

	reWhitespace := regexp.MustCompile(`\s+`)
	cleanedString = reWhitespace.ReplaceAllString(cleanedString, "")

	return cleanedString
}
