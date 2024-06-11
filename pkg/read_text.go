package pkg

import (
	"errors"

	"github.com/gildemberg-santos/webcrawlerurl_v2/util/extract"
	"github.com/gildemberg-santos/webcrawlerurl_v2/util/load_page"
	"github.com/gildemberg-santos/webcrawlerurl_v2/util/timestamp"
)

type ReadText struct {
	Url string
}

type responseSuccessReadText struct {
	ReadText struct {
		Text string `json:"text"`
	} `json:"readtext"`
	Url        string  `json:"url"`
	Timestamp  float64 `json:"ts"`
	StatusCode int     `json:"status_code"`
}

type responseErroReadtext struct {
	Erro       string      `json:"erro"`
	ReadText   interface{} `json:"readtext"`
	Url        string      `json:"url"`
	Timestamp  float64     `json:"ts"`
	StatusCode int         `json:"status_code"`
}

func NewReadText(url string) ReadText {
	return ReadText{
		Url: url,
	}
}

func (c *ReadText) Call() (interface{}, error) {
	ts := timestamp.NewTimestamp().Start()

	if c.Url == "" {
		err := errors.New("url is empty")
		ts.End()
		responseErro := responseErroReadtext{
			Erro:       err.Error(),
			ReadText:   nil,
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
			ReadText:   nil,
			Url:        c.Url,
			Timestamp:  ts.GetTime(),
			StatusCode: page.StatusCode,
		}
		return responseErro, err
	}

	informatin := extract.NewText(page.Source)
	extractext := informatin.Call()

	ts.End()
	responseSuccess := responseSuccessReadText{
		Url:        c.Url,
		Timestamp:  ts.GetTime(),
		StatusCode: page.StatusCode,
	}
	responseSuccess.ReadText.Text = extractext.Text

	return responseSuccess, nil
}
