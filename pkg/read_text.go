package pkg

import (
	"errors"

	"github.com/PuerkitoBio/goquery"
	"github.com/gildemberg-santos/webcrawlerurl_v2/util/extract"
	"github.com/gildemberg-santos/webcrawlerurl_v2/util/load_page"
	"github.com/gildemberg-santos/webcrawlerurl_v2/util/normalize"
	"github.com/gildemberg-santos/webcrawlerurl_v2/util/timestamp"
)

type ReadText struct {
	Url          string
	Sources      *goquery.Document
	Data         DataReadText
	LoadPageFast bool
}
type DataReadText struct {
	Text           string          `json:"text"`
	TotalCaracters int64           `json:"total_characters"`
	Url            string          `json:"url"`
	RetailerItemID string          `json:"retailer_item_id"`
	MetaTag        extract.MetaTag `json:"meta_tag,omitempty"`
}

type ResponseReadtext struct {
	Failure   bool           `json:"failure,omitempty"`
	Success   bool           `json:"success,omitempty"`
	Message   string         `json:"message,omitempty"`
	Data      []DataReadText `json:"data,omitempty"`
	Timestamp float64        `json:"ts"`
}

func NewReadText(url string, source *goquery.Document, loadPageFast bool) ReadText {
	return ReadText{
		Url:          url,
		Sources:      source,
		LoadPageFast: loadPageFast,
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
		page = load_page.NewLoadPage(c.Url, c.LoadPageFast)
		err = page.Call()
	} else {
		page = load_page.LoadPage{
			Url:          c.Url,
			Source:       c.Sources,
			LoadPageFast: c.LoadPageFast,
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
	meta_tag := extract.NewMetaTag(page.Source).Call()
	retailerItemID := normalize.NewNormalizeUrl(c.Url).MD5()
	extracMetaTag := meta_tag.Call()

	data := DataReadText{
		Text:           extractext.Text,
		TotalCaracters: int64(len(extractext.Text)),
		Url:            c.Url,
		RetailerItemID: retailerItemID,
		MetaTag:        *extracMetaTag,
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
