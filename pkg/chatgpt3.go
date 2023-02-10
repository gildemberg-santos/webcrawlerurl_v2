package pkg

import (
	"errors"
)

type ChatGpt3 struct {
	Url string
}

type responseSuccess struct {
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

type responseErro struct {
	Erro       string  `json:"erro"`
	Timestamp  float64 `json:"ts"`
	StatusCode int     `json:"status_code"`
}

func (c *ChatGpt3) Call(message string) (interface{}, error) {
	ts := Timestamp{}
	ts.Start()

	if c.Url == "" {
		err := errors.New("url is empty")
		ts.End()
		responseErro := responseErro{
			Erro:       err.Error(),
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
		responseErro := responseErro{
			Erro:       err.Error(),
			Timestamp:  ts.GetTime(),
			StatusCode: pagina.StatusCode,
		}
		return responseErro, err
	}

	informatin := ExtractInformation{}
	informatin.Init(pagina.Source, pagina.Url, 3, 5, 30)
	informatin.Call()

	score := Score{}
	score.Init(pagina.Url, &informatin)
	score.Call()

	ts.End()
	responseSuccess := responseSuccess{
		Url:        informatin.Url,
		Timestamp:  ts.GetTime(),
		Scone:      score.GetScore(),
		StatusCode: pagina.StatusCode,
	}
	responseSuccess.ChatGpt.Title = informatin.MainTitle
	responseSuccess.ChatGpt.Paragraph = informatin.MainParagraph
	responseSuccess.ChatGpt.Description = informatin.MetaDescription

	return responseSuccess, nil
}
