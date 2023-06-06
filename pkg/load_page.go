package pkg

import (
	"errors"
	"log"
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

	client := &http.Client{}

	req, err := http.NewRequest("GET", l.Url, nil)
	if err != nil {
		l.StatusCode = 404
		log.Fatal("Erro ao criar a requisição -> ", err)
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		err = errors.New("error to send request -> " + err.Error())
		log.Fatal(err)
		l.StatusCode = resp.StatusCode
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		l.StatusCode = resp.StatusCode
		err = errors.New("found error in the page")
		log.Fatal(err)
		return
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		l.StatusCode = 500
		err = errors.New("error to read body request -> " + err.Error())
		log.Fatal(err)
		return
	}

	l.Source = doc

	l.removerElementos()
	l.StatusCode = resp.StatusCode
	return
}

func (l *LoadPage) removerElementos() {
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
	l.Source.Find(tag).Each(func(_ int, s *goquery.Selection) {
		if s.HasClass(css) {
			s.Remove()
		}
	})
	defer doneLoadPage.Done()
}
