package extract

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gildemberg-santos/webcrawlerurl_v2/util/normalize"
)

type Link struct {
	Sources *goquery.Document `json:"-"`
	Limit   int64             `json:"-"`
	Url     string            `json:"url"`
	OutUrls []string          `json:"urls"`
}

func NewLink(source *goquery.Document, baseUrl string, limit int64) Link {
	return Link{
		Sources: source,
		Url:     baseUrl,
		Limit:   limit,
	}
}

func (l *Link) Call() *Link {
	l.extractUrls()
	l.removeDuplicationUrl()
	return l
}

func (l *Link) extractUrls() {
	l.Sources.Find("a").Each(func(_ int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		uri := normalize.NewNormalizeUrl(href)
		uri.BaseUrl = l.Url
		url, err := uri.GetUrl()
		if err != nil {
			return
		}
		l.OutUrls = append(l.OutUrls, url)
	})
}

func (l *Link) removeDuplicationUrl() {
	occurred := map[string]bool{}
	result := []string{}

	for i := range l.OutUrls {
		if l.Limit > 0 && int64(len(result)) >= l.Limit {
			break
		}
		if !occurred[l.OutUrls[i]] {
			occurred[l.OutUrls[i]] = true
			result = append(result, l.OutUrls[i])
		}
	}
	l.OutUrls = result
}
