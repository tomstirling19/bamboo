package utils

import (
	"regexp"
)

func StripString(s string) string {
	re := regexp.MustCompile(`\\n|\\`)
	strippedString := re.ReplaceAllString(s, "")

	reWhitespace := regexp.MustCompile(`\s+`)
	strippedString = reWhitespace.ReplaceAllString(strippedString, "")

	return strippedString
}
