package googleshopping

import (
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	useragent "github.com/gildemberg-santos/webcrawlerurl_v2/util/user_agent"
)

type GoogleShopping struct {
	UrlLocation string
	MaxTimeout  int64
	Feed        struct {
		Entry []struct {
			ID struct {
				Value string `xml:",chardata"`
			} `xml:"id"`
			Title struct {
				Value string `xml:",chardata"`
			} `xml:"title"`
			Summary struct {
				Value string `xml:",chardata"`
			} `xml:"summary"`
			Link struct {
				Value string `xml:",chardata"`
			} `xml:"link"`
			ImageLink struct {
				Value string `xml:",chardata"`
			} `xml:"image_link"`
			Price struct {
				Value string `xml:",chardata"`
			} `xml:"price"`
			Availability struct {
				Value string `xml:",chardata"`
			} `xml:"availability"`
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
	var currentTime = time.Now()
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
		timeoutThreshold := float64(g.MaxTimeout) * 0.95
		if time.Since(currentTime).Seconds() > timeoutThreshold {
			log.Println("Timeout threshold reached")
			break
		}
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
					ID struct {
						Value string `xml:",chardata"`
					} `xml:"id"`
					Title struct {
						Value string `xml:",chardata"`
					} `xml:"title"`
					Summary struct {
						Value string `xml:",chardata"`
					} `xml:"summary"`
					Link struct {
						Value string `xml:",chardata"`
					} `xml:"link"`
					ImageLink struct {
						Value string `xml:",chardata"`
					} `xml:"image_link"`
					Price struct {
						Value string `xml:",chardata"`
					} `xml:"price"`
					Availability struct {
						Value string `xml:",chardata"`
					} `xml:"availability"`
				}
				if err := decoder.DecodeElement(&entry, &se); err != nil {
					log.Default().Println("Error decoding entry: ", err)
					return err
				}

				entry.ID.Value = strings.TrimSpace(strings.ReplaceAll(entry.ID.Value, "\n", ""))
				entry.Title.Value = strings.TrimSpace(strings.ReplaceAll(entry.Title.Value, "\n", ""))
				entry.Summary.Value = strings.TrimSpace(strings.ReplaceAll(entry.Summary.Value, "\n", ""))
				entry.Link.Value = strings.TrimSpace(strings.ReplaceAll(entry.Link.Value, "\n", ""))
				entry.ImageLink.Value = strings.TrimSpace(strings.ReplaceAll(entry.ImageLink.Value, "\n", ""))
				entry.Price.Value = strings.TrimSpace(strings.ReplaceAll(entry.Price.Value, "\n", ""))
				entry.Availability.Value = strings.TrimSpace(strings.ReplaceAll(entry.Availability.Value, "\n", ""))

				g.Feed.Entry = append(g.Feed.Entry, entry)
			}
		}
	}

	return nil
}
