package pkg

import (
	"log"

	googleshopping "github.com/gildemberg-santos/webcrawlerurl_v2/util/google_shopping"
	"github.com/gildemberg-santos/webcrawlerurl_v2/util/url_match"
)

type EcommerceGoogleShopping struct {
	Visited    map[string]bool     `json:"-"`
	UrlPattern *url_match.UrlMatch `json:"-"`
	MaxTimeout int64               `json:"-"`
	Url        string              `json:"url,omitempty"`
	Urls       []string            `json:"urls,omitempty"`
}

func NewEcommerceGoogleShopping(url, urlPattern string, maxTimeout int64) *EcommerceGoogleShopping {
	return &EcommerceGoogleShopping{
		Url:        url,
		UrlPattern: url_match.NewUrlMatch(urlPattern),
		Visited:    map[string]bool{},
		MaxTimeout: maxTimeout,
	}
}

func (s *EcommerceGoogleShopping) Call() *EcommerceGoogleShopping {
	if err := s.crawler(s.Url); err != nil {
		log.Default().Println(err)
		return s
	}

	return s
}

func (s *EcommerceGoogleShopping) crawler(url string) error {
	s.Visited[url] = true

	googleShopping := googleshopping.NewGoogleShopping(url, s.MaxTimeout)

	if err := googleShopping.Call(); err != nil {
		log.Default().Println("Error request google shopping: ", err)
		return err
	}

	for _, entry := range googleShopping.Feed.Entry {
		if s.Visited[entry.Link.Value] {
			continue
		}
		if s.UrlPattern.Call(entry.Link.Value) {
			log.Println("Add url -> ", entry.Link.Value)
			s.Urls = append(s.Urls, entry.Link.Value)
		}
		s.Visited[entry.Link.Value] = true
	}

	return nil
}
