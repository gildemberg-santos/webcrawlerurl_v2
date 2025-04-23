package pkg

import (
	"errors"

	"github.com/gildemberg-santos/webcrawlerurl_v2/util/extract"
	"github.com/gildemberg-santos/webcrawlerurl_v2/util/load_page"
	"github.com/gildemberg-santos/webcrawlerurl_v2/util/timestamp"
)

type SmartCall struct {
	Url          string
	LoadPageFast bool
}

type responseSuccessGpt struct {
	SmartCall struct {
		Title       string `json:"main_header"`
		Paragraph   string `json:"main_paragraph"`
		Description string `json:"meta_description"`
	} `json:"chatgpt"`
	Url        string  `json:"url"`
	Timestamp  float64 `json:"ts"`
	Scone      float32 `json:"scone"`
	StatusCode int     `json:"status_code"`
}

type responseErroGpt struct {
	Erro       string      `json:"erro"`
	SmartCall  interface{} `json:"chatgpt"`
	Url        string      `json:"url"`
	Timestamp  float64     `json:"ts"`
	Scone      float32     `json:"scone"`
	StatusCode int         `json:"status_code"`
}

func NewSmartCall(url string, loadPageFast bool) SmartCall {
	return SmartCall{
		Url:          url,
		LoadPageFast: loadPageFast,
	}
}

func (c *SmartCall) Call() (interface{}, error) {
	ts := timestamp.NewTimestamp().Start()

	if c.Url == "" {
		err := errors.New("url is empty")
		ts.End()
		responseErro := responseErroGpt{
			Erro:       err.Error(),
			SmartCall:  nil,
			Url:        c.Url,
			Timestamp:  ts.GetTime(),
			Scone:      0,
			StatusCode: 500,
		}
		return responseErro, err
	}

	page := load_page.NewLoadPage(c.Url, c.LoadPageFast)

	err := page.Call()
	if err != nil {
		ts.End()
		responseErro := responseErroGpt{
			Erro:       err.Error(),
			SmartCall:  nil,
			Url:        c.Url,
			Timestamp:  ts.GetTime(),
			Scone:      0,
			StatusCode: page.StatusCode,
		}
		return responseErro, err
	}

	information := extract.NewLeadsterCustom(page.Source, 5, 5, 30)
	information.Call()

	score := Score{}
	score.Init(&information)
	score.Call()

	ts.End()
	responseSuccess := responseSuccessGpt{
		Url:        c.Url,
		Timestamp:  ts.GetTime(),
		Scone:      score.GetScore(),
		StatusCode: page.StatusCode,
	}
	responseSuccess.SmartCall.Title = information.TitleWebSite
	responseSuccess.SmartCall.Paragraph = information.MostRelevantText
	responseSuccess.SmartCall.Description = information.MetaDescription

	return responseSuccess, nil
}
