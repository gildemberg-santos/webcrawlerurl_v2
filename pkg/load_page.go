package pkg

import (
	"errors"
	"log"
	"net/http"
	"sync"
	"time"

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

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("GET", l.Url, nil)
	if err != nil {
		l.StatusCode = 404
		log.Default().Println("Error to create request -> ", err.Error())
		return
	}

	userAgentRandom := UserAgentRandom{}
	userAgentRandom.Call()
	log.Default().Println("User-Agent -> ", userAgentRandom.UserAgent)
	req.Header.Set("User-Agent", userAgentRandom.UserAgent)

	resp, err := client.Do(req)
	if err != nil {
		l.StatusCode = 404
		err = errors.New("error to send request -> " + err.Error())
		log.Default().Println(err.Error())
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		l.StatusCode = resp.StatusCode
		err = errors.New("found error in the page")
		log.Default().Println("Error to load page -> ", err.Error())
		return
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		l.StatusCode = 500
		err = errors.New("error to read body request -> " + err.Error())
		log.Default().Println(err.Error())
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
