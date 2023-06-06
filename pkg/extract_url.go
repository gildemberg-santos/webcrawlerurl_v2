package pkg

import (
	"github.com/PuerkitoBio/goquery"
)

type ExtractUrl struct {
	Url     string
	OutUrls []string
	Source  *goquery.Document
	Limit   int
}

func (e *ExtractUrl) Init(source *goquery.Document) {
	e.Source = source
}

func (e *ExtractUrl) Call() {
	e.extractUrls()
	e.removeDuplicationUrl()
}

func (e *ExtractUrl) extractUrls() {
	e.Source.Find("a").Each(func(_ int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		uri := NormalizeUrl{Url: href, BaseUrl: e.Url}
		url, err := uri.GetUrl()
		if err != nil {
			return
		}
		e.OutUrls = append(e.OutUrls, url)
	})
}

func (e *ExtractUrl) removeDuplicationUrl() {
	occurred := map[string]bool{}
	result := []string{}

	for i := range e.OutUrls {
		if e.Limit > 0 && len(result) >= e.Limit {
			break
		}
		if !occurred[e.OutUrls[i]] {
			occurred[e.OutUrls[i]] = true
			result = append(result, e.OutUrls[i])
		}
	}
	e.OutUrls = result
}
