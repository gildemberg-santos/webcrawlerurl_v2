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
type DataReadText struct {
	Text        string   `json:"text"`
	CountChunck int64    `json:"count_chunck"`
	Chuncks     []string `json:"chuncks"`
	Url         string   `json:"url"`
}

type responseReadtext struct {
	Failure    bool           `json:"failure"`
	Success    bool           `json:"success"`
	Message    string         `json:"message"`
	Data       []DataReadText `json:"data"`
	Timestamp  float64        `json:"ts"`
	StatusCode int            `json:"status_code"`
}

func NewReadText(url string, limit int64) ReadText {
	log.Println("NewReadText", url, limit)
	return ReadText{
		Url:   url,
		Limit: limit,
	}
}

func (c *ReadText) Call() (responseReadtext, error) {
	ts := timestamp.NewTimestamp().Start()

	if c.Url == "" {
		err := errors.New("url is empty")
		ts.End()
		responseErro := responseReadtext{
			Failure:    true,
			Success:    false,
			Message:    err.Error(),
			Timestamp:  ts.GetTime(),
			StatusCode: 500,
		}
		return responseErro, err
	}

	page := load_page.NewLoadPage(c.Url)

	err := page.Call()
	if err != nil {
		ts.End()
		responseErro := responseReadtext{
			Failure:    true,
			Success:    false,
			Message:    err.Error(),
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

	data := DataReadText{
		Text:        extractext.Text,
		CountChunck: chuncks.CountChunck,
		Chuncks:     chuncks.ListChuncks,
		Url:         c.Url,
	}
	datas := []DataReadText{data}

	responseSuccess := responseReadtext{
		Failure:    false,
		Success:    true,
		Message:    "Success",
		Data:       datas,
		Timestamp:  ts.GetTime(),
		StatusCode: page.StatusCode,
	}

	return responseSuccess, nil
}
