package pkg

import (
	"github.com/gildemberg-santos/webcrawlerurl_v2/util/load_page"
	"github.com/gildemberg-santos/webcrawlerurl_v2/util/normalize"
	"github.com/gildemberg-santos/webcrawlerurl_v2/util/timestamp"
)

type Ecommerce struct {
	Visited        map[string]bool     `json:"-"`
	MaxTimeout     int64               `json:"-"`
	IsLoadFast     bool                `json:"-"`
	WithTimestamp  timestamp.Timestamp `json:"-"`
	Urls           []string            `json:"urls,omitempty"`
	TotalCaracters int64               `json:"total_characters,omitempty"`
	Data           []DataReadText      `json:"data,omitempty"`
	Timestamp      float64             `json:"ts,omitempty"`
}

func NewEcommerce(urls []string, maxTimeout int64, isLoadFast bool) *Ecommerce {
	return &Ecommerce{
		Urls:          urls,
		MaxTimeout:    maxTimeout,
		IsLoadFast:    isLoadFast,
		Visited:       make(map[string]bool),
		WithTimestamp: *timestamp.NewTimestamp(),
	}
}

func (e *Ecommerce) Call() *Ecommerce {
	e.WithTimestamp.Start()
	for _, url := range e.Urls {
		e.crawler(url)
	}
	e.WithTimestamp.End()
	e.Timestamp = e.WithTimestamp.GetTime()

	return e
}

func (e *Ecommerce) crawler(url string) error {
	e.WithTimestamp.End()
	if e.WithTimestamp.GetTime() >= float64(e.MaxTimeout-5) {
		return nil
	}

	url, _ = normalize.NewNormalizeUrl(url).GetUrl()

	if e.Visited[url] {
		return nil
	}

	e.Visited[url] = true

	page := load_page.NewLoadPage(url, e.IsLoadFast)
	page.Timeout = 5
	page.Call()

	readText := NewReadText(url, page.Source, e.IsLoadFast)
	readText.Call()

	totalCaracters := readText.Data.TotalCaracters
	if totalCaracters > 0 {
		e.TotalCaracters += totalCaracters
		e.Data = append(e.Data, readText.Data)
	}

	return nil
}
