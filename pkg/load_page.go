package pkg

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type LoadPage struct {
	Url    string
	Source *goquery.Document
}

func (l *LoadPage) Load() error {
	res, err := http.Get(l.Url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return err
	}

	l.Source = doc

	l.RemoverElementos()
	return nil
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
