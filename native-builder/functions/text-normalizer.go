package functions

import (
	"regexp"
	"strings"
)

var removeWhitespaceRegexp = regexp.MustCompile(`\s+`)

func TextNormalizer(text string) string {
	text = removeWhitespaceRegexp.ReplaceAllString(text, "-")
	text = strings.ToLower(text)

	return text
}
