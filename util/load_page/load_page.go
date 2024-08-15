package load_page

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"github.com/gildemberg-santos/webcrawlerurl_v2/util/normalize"
	useragent "github.com/gildemberg-santos/webcrawlerurl_v2/util/user_agent"
)

var doneLoadPage sync.WaitGroup

type LoadPage struct {
	Url          string
	Timeout      time.Duration
	Source       *goquery.Document
	StatusCode   int
	LoadPageFast bool
}

func NewLoadPage(url string, loadPageFast bool) LoadPage {
	return LoadPage{
		Url:          url,
		Timeout:      10,
		LoadPageFast: loadPageFast,
	}
}

func (l *LoadPage) Call() (err error) {
	log.Default().Println("Loading page -> ", l.Url)
	if l.LoadPageFast {
		return l.loadPageFast()
	}

	return l.loadPageSlow()
}

func (l *LoadPage) loadPageFast() (err error) {
	_, err = normalize.NewNormalizeUrl(l.Url).GetUrl()

	if err != nil {
		l.StatusCode = 500
		return
	}

	client := &http.Client{
		Timeout: l.Timeout * time.Second,
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
		err = errors.New("Error to send request -> " + err.Error())
		log.Default().Println(err.Error())
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		l.StatusCode = resp.StatusCode
		err = fmt.Errorf("found error in the page status code -> %d", resp.StatusCode)
		log.Default().Println("Error to load page -> ", err.Error())
		return
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		l.StatusCode = 500
		err = errors.New("Error to read body request -> " + err.Error())
		log.Default().Println(err.Error())
		return
	}

	l.Source = doc

	l.removerElementos()
	l.StatusCode = resp.StatusCode
	return
}

func (l *LoadPage) loadPageSlow() (err error) {
	_, err = normalize.NewNormalizeUrl(l.Url).GetUrl()

	if err != nil {
		l.StatusCode = 500
		return
	}

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("no-sandbox", true),
		chromedp.Flag("remote-debugging-port", "9222"),
		chromedp.Flag("log-level", "0"),
		chromedp.Flag("v", "1"),
	)

	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx, chromedp.WithLogf(log.Printf))
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	var resp string

	err = chromedp.Run(ctx,
		chromedp.Navigate(l.Url),
		chromedp.Sleep(l.Timeout*time.Second),
		chromedp.OuterHTML(`html`, &resp, chromedp.ByQuery),
	)

	if err != nil {
		log.Default().Println("Error to load page -> ", err.Error())
		l.StatusCode = 500
		return
	}

	doc, err := goquery.NewDocumentFromReader(io.NopCloser(bytes.NewReader([]byte(resp))))
	if err != nil {
		l.StatusCode = 500
		err = errors.New("Error to read body request -> " + err.Error())
		log.Default().Println(err.Error())
		return
	}

	l.Source = doc

	l.removerElementos()
	l.StatusCode = 200
	return
}

func (l *LoadPage) removerElementos() {
	doneLoadPage.Add(12)
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
	go l.removeElementsDisplayNone("style", "")
	doneLoadPage.Wait()
}

func (l *LoadPage) removeElementsDisplayNone(tag string, css string) {
	l.Source.Find(tag).Each(func(_ int, s *goquery.Selection) {
		if s.HasClass(css) || (tag == "script" || tag == "noscript" || tag == "style") {
			s.Remove()
		}
	})
	defer doneLoadPage.Done()
}
