package googleshopping

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"

	useragent "github.com/gildemberg-santos/webcrawlerurl_v2/util/user_agent"
)

type GoogleShopping struct {
	UrlLocation string
	Feed        struct {
		Entry []struct {
			Link struct {
				Value string `xml:",chardata"`
			} `xml:"link"`
		} `xml:"entry"`
	} `xml:"feed"`
}

func NewGoogleShopping(url string) *GoogleShopping {
	return &GoogleShopping{
		UrlLocation: url,
	}
}

func (g *GoogleShopping) Call() error {
	if err := g.load(); err != nil {
		return err
	}

	return nil
}

func (g *GoogleShopping) load() error {
	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("GET", g.UrlLocation, nil)
	if err != nil {
		return err
	}

	req.Header.Set("User-Agent", useragent.NewUserAgentRandom().Call().UserAgent)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := xml.Unmarshal(body, &g.Feed); err != nil {
		return err
	}

	return nil
}
