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
	pattern := regexp.QuoteMeta(f.Pattern)
	pattern = strings.Replace(pattern, `\*\*`, `.*`, -1)
	pattern = "^" + pattern + "$"
	re := regexp.MustCompile(pattern)
	return re.MatchString(url)
}
