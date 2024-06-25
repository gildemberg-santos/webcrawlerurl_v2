package pkg

import (
	"errors"

	"github.com/PuerkitoBio/goquery"
	"github.com/gildemberg-santos/webcrawlerurl_v2/util/extract"
	"github.com/gildemberg-santos/webcrawlerurl_v2/util/load_page"
	"github.com/gildemberg-santos/webcrawlerurl_v2/util/timestamp"
)

type MappingUrl struct {
	Url     string
	Limit   int64
	Sources *goquery.Document
	Urls    []string
}

type ResponseUrls struct {
	Failure   bool     `json:"failure"`
	Success   bool     `json:"success"`
	Message   string   `json:"message"`
	Data      []string `json:"data"`
	Timestamp float64  `json:"ts"`
}

func NewMappingUrl(url string, limit int64, source *goquery.Document) MappingUrl {
	return MappingUrl{
		Url:     url,
		Limit:   limit,
		Sources: source,
	}
}

func (m *MappingUrl) Call() (ResponseUrls, error) {
	ts := timestamp.NewTimestamp().Start()

	if m.Url == "" {
		err := errors.New("url is empty")
		ts.End()
		responseErro := ResponseUrls{
			Failure:   true,
			Success:   false,
			Message:   err.Error(),
			Timestamp: ts.GetTime(),
		}
		return responseErro, err
	}

	var page load_page.LoadPage
	var err error

	if m.Sources == nil {
		page = load_page.NewLoadPage(m.Url)
		err = page.Call()
	} else {
		page = load_page.LoadPage{
			Url:    m.Url,
			Source: m.Sources,
		}
		err = nil
	}

	if err != nil {
		ts.End()
		responseErro := ResponseUrls{
			Failure:   true,
			Success:   false,
			Message:   err.Error(),
			Timestamp: ts.GetTime(),
		}
		return responseErro, err
	}

	extract_url := extract.NewLink(page.Source, m.Url, m.Limit)
	extract_url.Call()
	ts.End()

	responseSuccess := ResponseUrls{
		Failure:   false,
		Success:   true,
		Message:   "Success",
		Data:      extract_url.OutUrls,
		Timestamp: ts.GetTime(),
	}

	m.Urls = extract_url.OutUrls

	return responseSuccess, nil
}
