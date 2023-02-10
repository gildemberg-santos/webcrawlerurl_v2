package pkg

import (
	"errors"
	"net"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type LoadPage struct {
	Url        string
	Source     *goquery.Document
	StatusCode int
}

func (l *LoadPage) Load() (err error) {
	l.normalizeUrl()

	if !l.isUrl() {
		l.StatusCode = 500
		err = errors.New("Url is invalid")
		return
	}

	res, err := http.Get(l.Url)
	if err != nil {
		l.StatusCode = res.StatusCode
		return
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		l.StatusCode = res.StatusCode
		err = errors.New("Found error in the page")
		return
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		l.StatusCode = 500
		return
	}

	l.Source = doc

	l.RemoverElementos()
	l.StatusCode = res.StatusCode
	return
}

func (l *LoadPage) RemoverElementos() {
	l.removeElementsDisplayNone("div", "d-none")
	l.removeElementsDisplayNone("h1", "d-none")
	l.removeElementsDisplayNone("h2", "d-none")
	l.removeElementsDisplayNone("h3", "d-none")
	l.removeElementsDisplayNone("h4", "d-none")
	l.removeElementsDisplayNone("h5", "d-none")
	l.removeElementsDisplayNone("p", "d-none")
	l.removeElementsDisplayNone("span", "d-none")
	l.removeElementsDisplayNone("a", "d-none")
}

func (l *LoadPage) removeElementsDisplayNone(tag string, css string) {
	l.Source.Find(tag).Each(func(i int, s *goquery.Selection) {
		if s.HasClass(css) {
			s.Remove()
		}
	})
}

func (l *LoadPage) normalizeUrl() {
	if !strings.HasPrefix(l.Url, "https://") && !strings.HasPrefix(l.Url, "http://") {
		l.Url = "https://" + l.Url
	} else if strings.HasPrefix(l.Url, "http://") {
		l.Url = strings.Replace(l.Url, "http://", "https://", 1)
	}
}

func (l *LoadPage) isUrl() bool {
	url, err := url.ParseRequestURI(l.Url)
	if err != nil {
		return false
	}

	address := net.ParseIP(url.Host)
	if address == nil {
		return strings.Contains(url.Host, ".")
	}

	return true
}
