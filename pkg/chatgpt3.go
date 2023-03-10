package pkg

import (
	"errors"
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
	Erro       string  `json:"erro"`
	Url        string  `json:"url"`
	Timestamp  float64 `json:"ts"`
	StatusCode int     `json:"status_code"`
}

func (c *ChatGpt3) Call() (interface{}, error) {
	ts := Timestamp{}
	ts.Start()

	if c.Url == "" {
		err := errors.New("url is empty")
		ts.End()
		responseErro := responseErroGpt{
			Erro:       err.Error(),
			Url:        c.Url,
			Timestamp:  ts.GetTime(),
			StatusCode: 500,
		}
		return responseErro, err
	}

	pagina := LoadPage{
		Url: c.Url,
	}

	err := pagina.Load()
	if err != nil {
		ts.End()
		responseErro := responseErroGpt{
			Erro:       err.Error(),
			Url:        c.Url,
			Timestamp:  ts.GetTime(),
			StatusCode: pagina.StatusCode,
		}
		return responseErro, err
	}

	informatin := ExtractInformation{}
	informatin.Init(pagina.Source, 3, 5, 30)
	informatin.Call()

	score := Score{}
	score.Init(&informatin)
	score.Call()

	ts.End()
	responseSuccess := responseSuccessGpt{
		Url:        c.Url,
		Timestamp:  ts.GetTime(),
		Scone:      score.GetScore(),
		StatusCode: pagina.StatusCode,
	}
	responseSuccess.ChatGpt.Title = informatin.MainTitle
	responseSuccess.ChatGpt.Paragraph = informatin.MainParagraph
	responseSuccess.ChatGpt.Description = informatin.MetaDescription

	return responseSuccess, nil
}
