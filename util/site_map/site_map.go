package sitemap

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"

	useragent "github.com/gildemberg-santos/webcrawlerurl_v2/util/user_agent"
	"golang.org/x/net/html/charset"
)

type Link struct {
	Rel      string `xml:"rel,attr"`
	Hreflang string `xml:"hreflang,attr"`
	Href     string `xml:"href,attr"`
}

type URLs struct {
	Loc      string  `xml:"loc"`
	Lastmod  string  `xml:"lastmod"`
	Priority float32 `xml:"priority"`
	Link     []Link  `xml:"link"`
}

type Urlset struct {
	XMLName xml.Name `xml:"urlset"`
	URLs    []URLs   `xml:"url"`
}

type SitemapItem struct {
	Loc     string `xml:"loc"`
	Lastmod string `xml:"lastmod"`
}

type Sitemapindex struct {
	Sitemap []SitemapItem `xml:"sitemap"`
}

type SiteMap struct {
	UrlLocation  string
	Urlset       Urlset       `xml:"urlset"`
	Sitemapindex Sitemapindex `xml:"sitemapindex"`
}

func NewSiteMap(url string) *SiteMap {
	return &SiteMap{
		UrlLocation: url,
	}
}

func (s *SiteMap) Call() error {
	if err := s.load(); err != nil {
		return err
	}

	return nil
}

func (s *SiteMap) load() error {
	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("GET", s.UrlLocation, nil)
	if err != nil {
		return err
	}

	req.Header.Set("User-Agent", useragent.NewUserAgentRandom().Call().UserAgent)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return err
	}

	decoder := xml.NewDecoder(resp.Body)
	decoder.CharsetReader = charset.NewReaderLabel

	for {
		t, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "urlset" {
				var urlset Urlset
				if err := decoder.DecodeElement(&urlset, &se); err != nil {
					return err
				}

				s.Urlset = urlset
			}
			if se.Name.Local == "sitemapindex" {
				var sitemapindex Sitemapindex
				if err := decoder.DecodeElement(&sitemapindex, &se); err != nil {
					return err
				}

				s.Sitemapindex = sitemapindex
			}
		}
	}

	return nil
}
