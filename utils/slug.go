package utils

import (
	"regexp"
	"strings"
)

var slugRegex = regexp.MustCompile(`[^\w]+`)

func GenerateSlug(input string) string {
	slug := strings.ToLower(input)
	slug = slugRegex.ReplaceAllString(slug, "-")
	slug = strings.Trim(slug, "-")
	return slug
}
