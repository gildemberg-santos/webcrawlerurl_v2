package pkg

import (
	"errors"

	"github.com/gildemberg-santos/webcrawlerurl_v2/util/extract"
	"github.com/gildemberg-santos/webcrawlerurl_v2/util/load_page"
	"github.com/gildemberg-santos/webcrawlerurl_v2/util/timestamp"
)

type ChatGpt3 struct {
	Url string
}

type responseSuccessGpt struct {
	ChatGpt struct {
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
	ChatGpt3   interface{} `json:"chatgpt"`
	Url        string      `json:"url"`
	Timestamp  float64     `json:"ts"`
	Scone      float32     `json:"scone"`
	StatusCode int         `json:"status_code"`
}

func NewChatGpt3(url string) ChatGpt3 {
	return ChatGpt3{
		Url: url,
	}
}

func (c *ChatGpt3) Call() (interface{}, error) {
	ts := timestamp.NewTimestamp().Start()

	if c.Url == "" {
		err := errors.New("url is empty")
		ts.End()
		responseErro := responseErroGpt{
			Erro:       err.Error(),
			ChatGpt3:   nil,
			Url:        c.Url,
			Timestamp:  ts.GetTime(),
			Scone:      0,
			StatusCode: 500,
		}
		return responseErro, err
	}

	page := load_page.NewLoadPage(c.Url)

	err := page.Call()
	if err != nil {
		ts.End()
		responseErro := responseErroGpt{
			Erro:       err.Error(),
			ChatGpt3:   nil,
			Url:        c.Url,
			Timestamp:  ts.GetTime(),
			Scone:      0,
			StatusCode: page.StatusCode,
		}
		return responseErro, err
	}

	informatin := extract.NewLeadsterCustom(page.Source, 5, 5, 30)
	informatin.Call()

	score := Score{}
	score.Init(&informatin)
	score.Call()

	ts.End()
	responseSuccess := responseSuccessGpt{
		Url:        c.Url,
		Timestamp:  ts.GetTime(),
		Scone:      score.GetScore(),
		StatusCode: page.StatusCode,
	}
	responseSuccess.ChatGpt.Title = informatin.TitleWebSite
	responseSuccess.ChatGpt.Paragraph = informatin.MostRelevantText
	responseSuccess.ChatGpt.Description = informatin.MetaDescription

	return responseSuccess, nil
}
