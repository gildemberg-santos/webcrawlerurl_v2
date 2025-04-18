package load_page

import (
	"bytes"
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"github.com/gildemberg-santos/webcrawlerurl_v2/util/normalize"
	useragent "github.com/gildemberg-santos/webcrawlerurl_v2/util/user_agent"
)

type LoadPage struct {
	Url           string
	Timeout       time.Duration
	Source        *goquery.Document
	StatusCode    int
	LoadPageFast  bool
	SkipTLSVerify bool
}

func NewLoadPage(url string, loadPageFast bool) LoadPage {
	return LoadPage{
		Url:          url,
		Timeout:      10,
		LoadPageFast: loadPageFast,
	}
}

func (l *LoadPage) Call() error {
	log.Default().Println("Loading page -> ", l.Url)

	if l.LoadPageFast {
		if err := l.tryLoadPageFast(); err != nil {
			return err
		}
		return nil
	}

	return l.loadPageSlow()
}

func (l *LoadPage) tryLoadPageFast() error {
	if err := l.loadPageFast(); err != nil {
		if strings.Contains(err.Error(), "x509: certificate signed by unknown authority") {
			l.SkipTLSVerify = true
			return l.loadPageFast()
		}
		return err
	}
	return nil
}

func (l *LoadPage) loadPageFast() (err error) {
	_, err = normalize.NewNormalizeUrl(l.Url).GetUrl()

	if err != nil {
		l.StatusCode = 500
		return
	}

	client := &http.Client{Timeout: l.Timeout * time.Second}

	if l.SkipTLSVerify {
		client.Transport = &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
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
	l.removeElementsDisplayNone("div", "d-none")
	l.removeElementsDisplayNone("h1", "d-none")
	l.removeElementsDisplayNone("h2", "d-none")
	l.removeElementsDisplayNone("h3", "d-none")
	l.removeElementsDisplayNone("h4", "d-none")
	l.removeElementsDisplayNone("h5", "d-none")
	l.removeElementsDisplayNone("p", "d-none")
	l.removeElementsDisplayNone("span", "d-none")
	l.removeElementsDisplayNone("a", "d-none")
	l.removeElementsDisplayNone("script", "")
	l.removeElementsDisplayNone("noscript", "")
	l.removeElementsDisplayNone("style", "")
}

func (l *LoadPage) removeElementsDisplayNone(tag string, css string) {
	l.Source.Find(tag).Each(func(_ int, s *goquery.Selection) {
		if s.HasClass(css) || (tag == "script" || tag == "noscript" || tag == "style") {
			s.Remove()
		}
	})
}
