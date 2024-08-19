package pkg

import (
	"log"

	sitemap "github.com/gildemberg-santos/webcrawlerurl_v2/util/site_map"
	"github.com/gildemberg-santos/webcrawlerurl_v2/util/url_match"
)

type EcommerceSitemap struct {
	Visited           map[string]bool     `json:"-"`
	UrlSiteMapPattern *url_match.UrlMatch `json:"-"`
	UrlPattern        *url_match.UrlMatch `json:"-"`
	Url               string              `json:"url,omitempty"`
	Urls              []string            `json:"urls,omitempty"`
	UrlSiteMap        []string            `json:"sitemaps,omitempty"`
}

func NewEcommerceSitemap(url, urlPattern, urlSiteMapPattern string) *EcommerceSitemap {
	return &EcommerceSitemap{
		Url:               url,
		UrlPattern:        url_match.NewUrlMatch(urlPattern),
		UrlSiteMapPattern: url_match.NewUrlMatch(urlSiteMapPattern),
		Visited:           map[string]bool{},
	}
}

func (s *EcommerceSitemap) Call() *EcommerceSitemap {
	if err := s.crawler(s.Url); err != nil {
		log.Default().Println(err)
		return s
	}

	for _, url := range s.UrlSiteMap {
		if err := s.crawler(url); err != nil {
			log.Default().Println(err)
		}
	}

	return s
}

func (s *EcommerceSitemap) crawler(url string) error {
	s.Visited[url] = true

	siteMap := sitemap.NewSiteMap(url)

	if err := siteMap.Call(); err != nil {
		log.Default().Println(err)
		return err
	}

	for _, url := range siteMap.Urlset.Urls {
		if s.UrlPattern.Call(url.Loc) {
			s.Urls = append(s.Urls, url.Loc)
		}
	}

	for _, url := range siteMap.Sitemapindex.Sitemap {
		if s.UrlSiteMapPattern.Call(url.Loc) {
			s.UrlSiteMap = append(s.UrlSiteMap, url.Loc)
		}
	}

	return nil
}
