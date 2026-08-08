package utils

import (
	"regexp"
	"strings"

	"github.com/gosimple/slug"
)

// GenerateSlug creates a URL-friendly slug from text
func GenerateSlug(text string) string {
	// Use gosimple/slug for basic slug generation
	s := slug.Make(text)

	// Additional cleanup to match Node.js implementation
	s = strings.ToLower(s)
	s = strings.TrimSpace(s)

	// Remove special characters except hyphens
	reg := regexp.MustCompile(`[^\w\s-]`)
	s = reg.ReplaceAllString(s, "")

	// Replace multiple spaces with single hyphen
	reg = regexp.MustCompile(`\s+`)
	s = reg.ReplaceAllString(s, "-")

	// Replace multiple hyphens with single hyphen
	reg = regexp.MustCompile(`-+`)
	s = reg.ReplaceAllString(s, "-")

	return s
}
