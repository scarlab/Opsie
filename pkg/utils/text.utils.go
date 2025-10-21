package utils

import (
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/gosimple/slug"
)

// textUtils provides helper functions for text formatting, slugs, and casing.
type textUtils struct{}

// NewTextUtils creates a new instance of TextUtils.
func NewTextUtils() *textUtils {
	return &textUtils{}
}

var Text = NewTextUtils()

// Slugify converts text into a URL-safe slug.
func (t *textUtils) Slugify(s string) string {
	return slug.Make(s)
}

// TitleCase capitalizes the first letter of each word (Unicode-safe).
func (t *textUtils) TitleCase(s string) string {
	if s == "" {
		return s
	}
	var result []rune
	capNext := true
	for _, r := range s {
		if unicode.IsSpace(r) || unicode.IsPunct(r) {
			capNext = true
			result = append(result, r)
			continue
		}
		if capNext {
			result = append(result, unicode.ToTitle(r))
			capNext = false
		} else {
			result = append(result, unicode.ToLower(r))
		}
	}
	return string(result)
}

// ToLower safely converts string to lowercase.
func (t *textUtils) ToLower(s string) string {
	return strings.ToLower(s)
}

// ToUpper safely converts string to uppercase.
func (t *textUtils) ToUpper(s string) string {
	return strings.ToUpper(s)
}

// Truncate shortens a string to n runes and adds "…" if truncated.
func (t *textUtils) Truncate(s string, n int) string {
	if utf8.RuneCountInString(s) <= n {
		return s
	}
	runes := []rune(s)
	return string(runes[:n]) + "…"
}
