package pkg

import (
	"errors"

	"github.com/PuerkitoBio/goquery"
	"github.com/gildemberg-santos/webcrawlerurl_v2/util/extract"
	"github.com/gildemberg-santos/webcrawlerurl_v2/util/load_page"
	"github.com/gildemberg-santos/webcrawlerurl_v2/util/timestamp"
)

type ReadText struct {
	Url         string
	MaxChunck   int64
	MaxCaracter int64
	Sources     *goquery.Document
	Data        DataReadText
}
type DataReadText struct {
	Text           string   `json:"text"`
	TotalCaracters int64    `json:"total_characters"`
	CountChunck    int64    `json:"-"`
	Chuncks        []string `json:"-"`
	Url            string   `json:"url"`
}

type ResponseReadtext struct {
	Failure   bool           `json:"failure"`
	Success   bool           `json:"success"`
	Message   string         `json:"message"`
	Data      []DataReadText `json:"data"`
	Timestamp float64        `json:"ts"`
}

func NewReadText(url string, maxChunck, maxCaracter int64, source *goquery.Document) ReadText {
	return ReadText{
		Url:         url,
		MaxChunck:   maxChunck,
		MaxCaracter: maxCaracter,
		Sources:     source,
	}
}

func (c *ReadText) Call() (ResponseReadtext, error) {
	ts := timestamp.NewTimestamp().Start()

	if c.Url == "" {
		err := errors.New("url is empty")
		ts.End()
		responseErro := ResponseReadtext{
			Failure:   true,
			Success:   false,
			Message:   err.Error(),
			Data:      []DataReadText{},
			Timestamp: ts.GetTime(),
		}
		return responseErro, err
	}

	var page load_page.LoadPage
	var err error

	if c.Sources == nil {
		page = load_page.NewLoadPage(c.Url)
		err = page.Call()
	} else {
		page = load_page.LoadPage{
			Url:    c.Url,
			Source: c.Sources,
		}
		err = nil
	}

	if err != nil {
		ts.End()
		responseErro := ResponseReadtext{
			Failure:   true,
			Success:   false,
			Message:   err.Error(),
			Data:      []DataReadText{},
			Timestamp: ts.GetTime(),
		}
		return responseErro, err
	}

	informatin := extract.NewText(page.Source)
	extractext := informatin.Call()

	data := DataReadText{
		Text:           extractext.Text,
		TotalCaracters: int64(len(extractext.Text)),
		Url:            c.Url,
	}
	c.Data = data
	datas := []DataReadText{data}

	ts.End()

	responseSuccess := ResponseReadtext{
		Failure:   false,
		Success:   true,
		Message:   "Success",
		Data:      datas,
		Timestamp: ts.GetTime(),
	}

	return responseSuccess, nil
}
