package pkg

import (
	"errors"

	"github.com/gildemberg-santos/webcrawlerurl_v2/util/extract"
	"github.com/gildemberg-santos/webcrawlerurl_v2/util/timestamp"
)

type MappingUrl struct {
	Url   string
	Limit int
}

type responseSuccessUrls struct {
	MappingUrl struct {
		Urls []string `json:"urls"`
	} `json:"mapping_url"`
	Url        string  `json:"url"`
	Timestamp  float64 `json:"ts"`
	StatusCode int     `json:"status_code"`
}

type responseErroUrls struct {
	Erro       string  `json:"erro"`
	Url        string  `json:"url"`
	Timestamp  float64 `json:"ts"`
	StatusCode int     `json:"status_code"`
}

func NewMappingUrl(url string, limit int) MappingUrl {
	return MappingUrl{
		Url:   url,
		Limit: limit,
	}
}

func (m *MappingUrl) Call() (interface{}, error) {
	ts := timestamp.NewTimestamp().Start()

	if m.Url == "" {
		err := errors.New("url is empty")
		ts.End()
		responseErro := responseErroUrls{
			Erro:       err.Error(),
			Url:        m.Url,
			Timestamp:  ts.GetTime(),
			StatusCode: 500,
		}
		return responseErro, err
	}

	page := NewLoadPage(m.Url)

	err := page.Call()
	if err != nil {
		ts.End()
		responseErro := responseErroUrls{
			Erro:       err.Error(),
			Url:        m.Url,
			Timestamp:  ts.GetTime(),
			StatusCode: page.StatusCode,
		}
		return responseErro, err
	}

	extract_url := extract.NewLink(page.Source, m.Url, m.Limit)
	extract_url.Call()
	ts.End()

	responseSuccess := responseSuccessUrls{
		Url:        m.Url,
		Timestamp:  ts.GetTime(),
		StatusCode: page.StatusCode,
	}
	responseSuccess.MappingUrl.Urls = extract_url.OutUrls

	return responseSuccess, nil
}
