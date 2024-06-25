package pkg

import (
	"log"

	"github.com/gildemberg-santos/webcrawlerurl_v2/util/load_page"
	"github.com/gildemberg-santos/webcrawlerurl_v2/util/normalize"
	"github.com/gildemberg-santos/webcrawlerurl_v2/util/timestamp"
)

type LeadsterAI struct {
	Visited          map[string]bool `json:"-"`
	Url              string          `json:"-"`
	MaxUrlLimit      int64           `json:"-"`
	MaxChunckLimit   int64           `json:"-"`
	MaxCaracterLimit int64           `json:"-"`
	CountChunck      int64           `json:"-"`
	TotalCaracters   int64           `json:"total_characters"`
	Data             []DataReadText  `json:"data"`
	Timestamp        float64         `json:"ts"`
}

func NewLeadsterAI(url string, maxUrlLimit int64, maxChunckLimit int64, maxCaracterLimit int64) LeadsterAI {
	return LeadsterAI{
		Url:              url,
		MaxUrlLimit:      maxUrlLimit,
		MaxChunckLimit:   maxChunckLimit,
		MaxCaracterLimit: maxCaracterLimit,
		Visited:          make(map[string]bool),
	}
}

func (l *LeadsterAI) Call() *LeadsterAI {
	log.Println("Start LeadsterAI")
	ts := timestamp.NewTimestamp().Start()
	if l.MaxUrlLimit > 0 {
		l.crawler(l.Url)
	}
	timestamp.NewTimestamp().End()
	ts.End()
	l.Timestamp = ts.GetTime()

	return l
}

func (l *LeadsterAI) crawler(url string) {
	url, _ = normalize.NewNormalizeUrl(url).GetUrl()
	if l.Visited[url] {
		return
	}

	if int64(len(l.Data)) >= l.MaxUrlLimit {
		return
	}

	l.Visited[url] = true

	page := load_page.NewLoadPage(url)
	page.Timeout = 5
	page.Call()

	mapping := NewMappingUrl(url, l.MaxUrlLimit, page.Source)
	if l.MaxUrlLimit > 1 {
		mapping.Call()
	}

	readText := NewReadText(url, l.MaxChunckLimit, l.MaxCaracterLimit, page.Source)
	readText.Call()

	if readText.Data.TotalCaracters > 0 {
		l.TotalCaracters += readText.Data.TotalCaracters
		l.Data = append(l.Data, readText.Data)
	}

	for _, tmp_url := range mapping.Urls {
		tmp_url, _ = normalize.NewNormalizeUrl(tmp_url).GetUrl()
		if l.Visited[tmp_url] {
			continue
		}
		l.crawler(tmp_url)
	}
}
