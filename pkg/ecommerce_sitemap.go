package pkg

import (
	"log"

	"github.com/gildemberg-santos/webcrawlerurl_v2/util/normalize"
	sitemap "github.com/gildemberg-santos/webcrawlerurl_v2/util/site_map"
	"github.com/gildemberg-santos/webcrawlerurl_v2/util/url_match"
)

type EcommerceSitemap struct {
	Visited           map[string]bool     `json:"-"`
	Url               string              `json:"url"`
	UrlPattern        *url_match.UrlMatch `json:"-"`
	UrlSiteMapPattern *url_match.UrlMatch `json:"-"`
	Urls              []string            `json:"urls"`
	UrlSiteMap        []string            `json:"sitemaps"`
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
		normalizeUrl, _ := normalize.NewNormalizeUrl(url.Loc).GetUrl()
		if s.UrlPattern.Call(normalizeUrl) {
			s.Urls = append(s.Urls, normalizeUrl)
		}
	}

	for _, url := range siteMap.Sitemapindex.Sitemap {
		if s.UrlSiteMapPattern.Call(url.Loc) {
			s.UrlSiteMap = append(s.UrlSiteMap, url.Loc)
		}
	}

	return nil
}