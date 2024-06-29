package site_map

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"

	useragent "github.com/gildemberg-santos/webcrawlerurl_v2/util/user_agent"
)

type SiteMap struct {
	UrlLocation string
	Urlset      struct {
		XMLName xml.Name `xml:"urlset"`
		Urls    []struct {
			Loc      string  `xml:"loc"`
			Lastmod  string  `xml:"lastmod"`
			Priority float32 `xml:"priority"`
		} `xml:"url"`
	}
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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = xml.Unmarshal(body, &s.Urlset)
	if err != nil {
		return err
	}

	return nil
}
