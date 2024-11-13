package googleshopping

import (
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"time"

	useragent "github.com/gildemberg-santos/webcrawlerurl_v2/util/user_agent"
	"golang.org/x/net/html/charset"
)

type GoogleShopping struct {
	UrlLocation string
	MaxTimeout  int64
	isRSS       bool
	Feed        Feed `xml:"feed"`
	RSS         RSS  `xml:"rss"`
	RDF         RDF  `xml:"rdf"`
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
	decoder.CharsetReader = charset.NewReaderLabel

	for {
		timeoutThreshold := float64(g.MaxTimeout) * 0.95
		if time.Since(currentTime).Seconds() > timeoutThreshold {
			log.Default().Println("Timeout threshold reached")
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
			if se.Name.Local == "rss" {
				g.isRSS = true
			}
			if se.Name.Local == "rdf" {
				g.isRSS = false
			}
			if se.Name.Local == "entry" {
				var entry Entry
				if err := decoder.DecodeElement(&entry, &se); err != nil {
					log.Default().Println("Error decoding entry: ", err)
					return err
				}

				g.Feed.AddEntry(entry)
			}
			if se.Name.Local == "item" {
				var item Item
				if err := decoder.DecodeElement(&item, &se); err != nil {
					log.Default().Println("Error decoding item: ", err)
					return err
				}

				if g.isRSS {
					g.RSS.AddItem(item)
				} else {
					g.RDF.AddItem(item)
				}

			}
		}
	}
	g.Feed.Normalize()
	g.RSS.Normalize()
	g.RDF.Normalize()

	return nil
}
