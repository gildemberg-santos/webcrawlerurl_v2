package pkg

import (
	"errors"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gildemberg-santos/webcrawlerurl_v2/util/normalize"
	useragent "github.com/gildemberg-santos/webcrawlerurl_v2/util/user_agent"
)

var doneLoadPage sync.WaitGroup

type LoadPage struct {
	Url        string
	Source     *goquery.Document
	StatusCode int
}

func (l *LoadPage) Load() (err error) {
	_, err = normalize.NewNormalizeUrl(l.Url).GetUrl()

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

	req.Header.Set("User-Agent", useragent.NewUserAgentRandom().Call().UserAgent)

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
	doneLoadPage.Add(11)
	go l.removeElementsDisplayNone("div", "d-none")
	go l.removeElementsDisplayNone("h1", "d-none")
	go l.removeElementsDisplayNone("h2", "d-none")
	go l.removeElementsDisplayNone("h3", "d-none")
	go l.removeElementsDisplayNone("h4", "d-none")
	go l.removeElementsDisplayNone("h5", "d-none")
	go l.removeElementsDisplayNone("p", "d-none")
	go l.removeElementsDisplayNone("span", "d-none")
	go l.removeElementsDisplayNone("a", "d-none")
	go l.removeElementsDisplayNone("script", "")
	go l.removeElementsDisplayNone("noscript", "")
	doneLoadPage.Wait()
}

func (l *LoadPage) removeElementsDisplayNone(tag string, css string) {
	l.Source.Find(tag).Each(func(_ int, s *goquery.Selection) {
		if s.HasClass(css) || (tag == "script" || tag == "noscript") {
			s.Remove()
		}
	})
	defer doneLoadPage.Done()
}
