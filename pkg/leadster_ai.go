package pkg

import (
	"strings"

	"github.com/gildemberg-santos/webcrawlerurl_v2/util/load_page"
	"github.com/gildemberg-santos/webcrawlerurl_v2/util/normalize"
	sitemap "github.com/gildemberg-santos/webcrawlerurl_v2/util/site_map"
	"github.com/gildemberg-santos/webcrawlerurl_v2/util/timestamp"
	"github.com/gildemberg-santos/webcrawlerurl_v2/util/url_match"
)

type LeadsterAI struct {
	Visited        map[string]bool     `json:"-"`
	Url            string              `json:"-"`
	MaxUrlLimit    int64               `json:"-"`
	MaxTimeout     int64               `json:"-"`
	FilterUrlMatch *url_match.UrlMatch `json:"-"`
	IsLoadFast     bool                `json:"-"`
	WithTimestamp  timestamp.Timestamp `json:"-"`
	TotalCaracters int64               `json:"total_characters"`
	Data           []DataReadText      `json:"data,omitempty"`
	Timestamp      float64             `json:"ts"`
}

func NewLeadsterAI(url, urlPattern string, maxUrlLimit, maxTimeout int64, isLoadFast bool, discardedUrls []string) LeadsterAI {
	visited := make(map[string]bool)

	for _, url := range discardedUrls {
		visited[url] = true
	}

	return LeadsterAI{
		Url:            url,
		MaxUrlLimit:    maxUrlLimit,
		Visited:        visited,
		FilterUrlMatch: url_match.NewUrlMatch(urlPattern),
		IsLoadFast:     isLoadFast,
		MaxTimeout:     maxTimeout,
		WithTimestamp:  *timestamp.NewTimestamp(),
	}
}

func (l *LeadsterAI) Call(isSiteMap, isComplete bool) *LeadsterAI {
	l.WithTimestamp.Start()
	if l.MaxUrlLimit > 0 {
		l.crawler(l.Url, isSiteMap, isComplete)
	}
	l.WithTimestamp.End()
	l.Timestamp = l.WithTimestamp.GetTime()

	return l
}

func (l *LeadsterAI) crawler(url string, isSiteMap, isComplete bool) {
	l.WithTimestamp.End()
	if l.WithTimestamp.GetTime() >= float64(l.MaxTimeout-5) {
		return
	}

	url, _ = normalize.NewNormalizeUrl(url).GetUrl()

	if l.Visited[url] {
		return
	}

	if int64(len(l.Data)) >= l.MaxUrlLimit {
		return
	}

	l.Visited[url] = true

	page := load_page.NewLoadPage(url, l.IsLoadFast)
	page.Timeout = 5
	page.Call()

	mapping := NewMappingUrl(url, l.MaxUrlLimit, l.IsLoadFast, page.Source)
	if l.MaxUrlLimit > 1 {
		mapping.Call()
	}

	readText := NewReadText(url, page.Source, l.IsLoadFast)
	readText.Call()

	matchUrl := l.FilterUrlMatch.Call(url)
	isExtensionsXML := !strings.Contains(url, ".xml")
	totalCaracters := readText.Data.TotalCaracters

	if totalCaracters > 0 && matchUrl && isExtensionsXML {
		l.TotalCaracters += totalCaracters
		l.Data = append(l.Data, readText.Data)
	}

	if isSiteMap {
		if !strings.Contains(url, ".xml") {
			url = url + "/sitemap.xml"
		}

		siteMap := sitemap.NewSiteMap(url)
		if err := siteMap.Call(); err == nil {
			for _, tmp_url := range siteMap.Sitemapindex.Sitemap {
				if l.Visited[tmp_url.Loc] {
					continue
				}
				l.crawler(tmp_url.Loc, true, false)
			}
			for _, tmp_url := range siteMap.Urlset.Urls {
				tmp_url.Loc, _ = normalize.NewNormalizeUrl(tmp_url.Loc).GetUrl()
				if l.Visited[tmp_url.Loc] {
					continue
				}
				l.crawler(tmp_url.Loc, false, isComplete)
			}
		}
	}

	if isComplete {
		for _, tmp_url := range mapping.Urls {
			tmp_url, _ = normalize.NewNormalizeUrl(tmp_url).GetUrl()
			if l.Visited[tmp_url] {
				continue
			}
			l.crawler(tmp_url, false, true)
		}
	}
}
