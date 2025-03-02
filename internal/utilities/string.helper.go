package utilities

import (
	"regexp"
	"strings"
)

// ToKebabCase converts a given string to kebab-case.
// It replaces spaces and underscores with hyphens, removes non-alphanumeric characters,
// and ensures there are no consecutive hyphens.
func ToKebabCase(s string) string {
	s = strings.ToLower(s)

	s = strings.ReplaceAll(s, " ", "-")
	s = strings.ReplaceAll(s, "_", "-")

	re := regexp.MustCompile(`[^a-z0-9]+`)
	s = re.ReplaceAllString(s, "-")

	s = regexp.MustCompile(`-+`).ReplaceAllString(s, "-")

	s = strings.Trim(s, "-")

	return s
}
