package utils

import (
	"regexp"
	"strings"
)

var removeWhitespaceRegexp = regexp.MustCompile(`\s+`)

func TextNormalizer(text string) string {
	text = strings.Trim(text, " ")
	text = removeWhitespaceRegexp.ReplaceAllString(text, "-")
	text = strings.ToLower(text)
	text = strings.Trim(text, " ")
	text = strings.Trim(text, "-")
	text = strings.Trim(text, "/")
	text = strings.Trim(text, ".")

	return text
}
