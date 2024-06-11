package pkg

import (
	"errors"
	"log"

	"github.com/gildemberg-santos/webcrawlerurl_v2/util/chunck"
	"github.com/gildemberg-santos/webcrawlerurl_v2/util/extract"
	"github.com/gildemberg-santos/webcrawlerurl_v2/util/load_page"
	"github.com/gildemberg-santos/webcrawlerurl_v2/util/timestamp"
)

type ReadText struct {
	Url   string
	Limit int64
}

type responseSuccessReadText struct {
	Text        string   `json:"text"`
	CountChunck int64    `json:"count_chunck"`
	Chuncks     []string `json:"chuncks"`
	Url         string   `json:"url"`
	Timestamp   float64  `json:"ts"`
	StatusCode  int      `json:"status_code"`
}

type responseErroReadtext struct {
	Erro       string  `json:"erro"`
	Url        string  `json:"url"`
	Timestamp  float64 `json:"ts"`
	StatusCode int     `json:"status_code"`
}

func NewReadText(url string, limit int64) ReadText {
	log.Println("NewReadText", url, limit)
	return ReadText{
		Url:   url,
		Limit: limit,
	}
}

func (c *ReadText) Call() (interface{}, error) {
	ts := timestamp.NewTimestamp().Start()

	if c.Url == "" {
		err := errors.New("url is empty")
		ts.End()
		responseErro := responseErroReadtext{
			Erro:       err.Error(),
			Url:        c.Url,
			Timestamp:  ts.GetTime(),
			StatusCode: 500,
		}
		return responseErro, err
	}

	page := load_page.NewLoadPage(c.Url)

	err := page.Call()
	if err != nil {
		ts.End()
		responseErro := responseErroReadtext{
			Erro:       err.Error(),
			Url:        c.Url,
			Timestamp:  ts.GetTime(),
			StatusCode: page.StatusCode,
		}
		return responseErro, err
	}

	informatin := extract.NewText(page.Source)
	extractext := informatin.Call()
	chuncks := chunck.NewChunck(extractext.Text, c.Limit)
	chuncks.Call()

	ts.End()
	responseSuccess := responseSuccessReadText{
		Text:        extractext.Text,
		CountChunck: chuncks.CountChunck,
		Chuncks:     chuncks.ListChuncks,
		Url:         c.Url,
		Timestamp:   ts.GetTime(),
		StatusCode:  page.StatusCode,
	}

	return responseSuccess, nil
}
