package utils

import (
	"regexp"
	"strings"
)

// Slugify converts a string into a slug suitable for URLs and database names.
func Slugify(s string) string {
	// Convert to lowercase
	s = strings.ToLower(s)
	// Replace spaces and underscores with hyphens
	s = strings.ReplaceAll(s, " ", "-")
	s = strings.ReplaceAll(s, "_", "-")
	// Remove all non-alphanumeric and non-hyphen characters
	re := regexp.MustCompile(`[^a-z0-9\-]`)
	s = re.ReplaceAllString(s, "")
	// Remove duplicate hyphens
	re2 := regexp.MustCompile(`-+`)
	s = re2.ReplaceAllString(s, "-")
	// Trim leading and trailing hyphens
	s = strings.Trim(s, "-")
	return s
}
