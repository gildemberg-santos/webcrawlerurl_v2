package pkg

import (
	"log"
	"sync"

	"github.com/gildemberg-santos/webcrawlerurl_v2/util/load_page"
	"github.com/gildemberg-santos/webcrawlerurl_v2/util/normalize"
	"github.com/gildemberg-santos/webcrawlerurl_v2/util/timestamp"
)

type Ecommerce struct {
	Visited        sync.Map            `json:"-"`
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
		Visited:       sync.Map{},
		WithTimestamp: *timestamp.NewTimestamp(),
	}
}

func (e *Ecommerce) Call() *Ecommerce {
	e.WithTimestamp.Start()

	limitThreads := make(chan struct{}, 2)
	done := make(chan bool, len(e.Urls))

	for _, url := range e.Urls {
		limitThreads <- struct{}{}

		go func(url string) {
			defer func() {
				done <- true
				<-limitThreads
			}()

			if err := e.crawler(url); err != nil {
				log.Printf("Erro ao processar URL %s: %v", url, err)
				return
			}
		}(url)
	}

	for i := 0; i < len(e.Urls); i++ {
		<-done
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

	if _, loaded := e.Visited.LoadOrStore(url, true); loaded {
		// Se o URL já está presente no mapa, retorne sem fazer nada
		return nil
	}

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
