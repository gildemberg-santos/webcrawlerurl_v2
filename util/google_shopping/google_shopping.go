package googleshopping

import (
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"time"

	useragent "github.com/gildemberg-santos/webcrawlerurl_v2/util/user_agent"
)

type GoogleShopping struct {
	UrlLocation string
	MaxTimeout  int64
	Feed        struct {
		Entry []struct {
			Link struct {
				Value string `xml:",chardata"`
			} `xml:"link"`
		} `xml:"entry"`
	} `xml:"feed"`
}

func NewGoogleShopping(url string, maxTimeout int64) *GoogleShopping {
	return &GoogleShopping{
		UrlLocation: url,
		MaxTimeout:  maxTimeout,
	}
}

func (g *GoogleShopping) Call() error {
	if err := g.load(); err != nil {
		return err
	}

	return nil
}

func (g *GoogleShopping) load() error {
	log.Println("Start crawler google shopping")
	client := &http.Client{Timeout: time.Duration(g.MaxTimeout) * time.Second}
	req, err := http.NewRequest("GET", g.UrlLocation, nil)
	if err != nil {
		log.Default().Println("Error request google shopping: ", err)
		return err
	}

	req.Header.Set("User-Agent", useragent.NewUserAgentRandom().Call().UserAgent)

	resp, err := client.Do(req)
	if err != nil {
		log.Default().Println("Error request google shopping: ", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("Status code: ", resp.StatusCode)
		return err
	}

	decoder := xml.NewDecoder(resp.Body)
	for {
		t, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Default().Println("Error decoding token: ", err)
			return err
		}

		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "entry" {
				var entry struct {
					Link struct {
						Value string `xml:",chardata"`
					} `xml:"link"`
				}
				if err := decoder.DecodeElement(&entry, &se); err != nil {
					log.Default().Println("Error decoding entry: ", err)
					return err
				}
				g.Feed.Entry = append(g.Feed.Entry, entry)
			}
		}
	}

	return nil
}
