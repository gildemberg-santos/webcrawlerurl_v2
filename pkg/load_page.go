package pkg

import (
	"errors"
	"net/http"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

var doneLoadPage sync.WaitGroup

type LoadPage struct {
	Url        string
	Source     *goquery.Document
	StatusCode int
}

func (l *LoadPage) Load() (err error) {
	uri := NormalizeUrl{Url: l.Url}
	_, err = uri.GetUrl()

	if err != nil {
		l.StatusCode = 500
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
	doneLoadPage.Add(9)
	go l.removeElementsDisplayNone("div", "d-none")
	go l.removeElementsDisplayNone("h1", "d-none")
	go l.removeElementsDisplayNone("h2", "d-none")
	go l.removeElementsDisplayNone("h3", "d-none")
	go l.removeElementsDisplayNone("h4", "d-none")
	go l.removeElementsDisplayNone("h5", "d-none")
	go l.removeElementsDisplayNone("p", "d-none")
	go l.removeElementsDisplayNone("span", "d-none")
	go l.removeElementsDisplayNone("a", "d-none")
	doneLoadPage.Wait()
}

func (l *LoadPage) removeElementsDisplayNone(tag string, css string) {
	l.Source.Find(tag).Each(func(i int, s *goquery.Selection) {
		if s.HasClass(css) {
			s.Remove()
		}
	})
	defer doneLoadPage.Done()
}
