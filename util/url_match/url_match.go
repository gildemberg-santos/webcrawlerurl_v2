package url_match

import (
	"regexp"
	"strings"
)

type UrlMatch struct {
	Url     string
	Pattern string
}

func NewUrlMatch(pattern string) *UrlMatch {
	return &UrlMatch{
		Pattern: pattern,
	}
}

func (f *UrlMatch) Call(url string) bool {
	if f.Pattern == "" {
		return true
	}

	pattern := strings.Replace(f.Pattern, "**", ".*", -1)

	pattern = strings.Replace(pattern, "*", "[^/]*", -1)

	re := regexp.MustCompile("^" + pattern + "$")

	return re.MatchString(url)
}
