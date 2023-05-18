package pkg

import (
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

var doneExtractInformation sync.WaitGroup

type ExtractInformation struct {
	TitleWebSite        string
	MostRelevantText    string
	MetaDescription     string
	TitleWebSiteMin     int
	MostRelevantTextMin int
	MetaDescriptionMin  int
	Source              *goquery.Document
}

func (e *ExtractInformation) Init(source *goquery.Document, titleWebSiteMin, mostRelevantTextMin, metaDescriptionMin int) {
	e.Source = source
	e.TitleWebSiteMin = titleWebSiteMin
	e.MostRelevantTextMin = mostRelevantTextMin
	e.MetaDescriptionMin = metaDescriptionMin
}

func (e *ExtractInformation) Call() {
	doneExtractInformation.Add(3)
	go e.extractTitleWebSite()
	go e.extractMostRelevantText()
	go e.extractMetaDescription()
	doneExtractInformation.Wait()

	e.normalize()
}

func (e *ExtractInformation) extractTitleWebSite() {
	e.Source.Find("head").Each(func(_ int, s *goquery.Selection) {
		s.Find("title").Each(func(_ int, s *goquery.Selection) {
			text := s.Text()
			text = strings.TrimSpace(text)
			e.TitleWebSite = text
		})
	})

	e.Source.Find("meta").Each(func(_ int, s *goquery.Selection) {
		name, _ := s.Attr("property")
		if name == "og:title" && e.TitleWebSite == "" {
			content, _ := s.Attr("content")
			content = strings.TrimSpace(content)
			e.TitleWebSite = content
		}

		if name == "twitter:title" && e.TitleWebSite == "" {
			content, _ := s.Attr("content")
			content = strings.TrimSpace(content)
			e.TitleWebSite = content
		}
	})
	defer doneExtractInformation.Done()
}

func (e *ExtractInformation) extractMostRelevantText() {
	e.filterMostRelevantText("h1")

	if e.MostRelevantText != "" {
		defer doneExtractInformation.Done()
		return
	}

	e.filterMostRelevantText("h2")
	defer doneExtractInformation.Done()
}

func (e *ExtractInformation) extractMetaDescription() {
	e.Source.Find("meta").Each(func(_ int, s *goquery.Selection) {
		name, _ := s.Attr("name")
		if name == "description" {
			content, _ := s.Attr("content")
			e.MetaDescription = content
		}

		if name == "og:description" && e.MetaDescription == "" {
			content, _ := s.Attr("content")
			e.MetaDescription = content
		}

		if name == "twitter:description" && e.MetaDescription == "" {
			content, _ := s.Attr("content")
			e.MetaDescription = content
		}

	})

	defer doneExtractInformation.Done()
}

func (e *ExtractInformation) filterMostRelevantText(tag string) {
	var title = make([]string, 0)
	var first string = ""

	e.Source.Find(tag).Each(func(_ int, s *goquery.Selection) {
		text := s.Text()
		text = strings.TrimSpace(text)

		if text != "" {
			if first == "" {
				first = text
			}
			title = append(title, text)
		}
	})

	for _, t := range title {
		words := strings.Split(t, " ")
		if len(words) >= e.MostRelevantTextMin && e.MostRelevantText == "" {
			e.MostRelevantText = t
		}
	}

	if e.MostRelevantText == "" {
		e.MostRelevantText = first
	}
}

func (e *ExtractInformation) normalize() {
	remove := []string{"\n", "\t", "\r"}

	for _, r := range remove {
		e.TitleWebSite = strings.Replace(e.TitleWebSite, r, " ", -1)
		e.TitleWebSite = strings.TrimSpace(e.TitleWebSite)

		e.MostRelevantText = strings.Replace(e.MostRelevantText, r, " ", -1)
		e.MostRelevantText = strings.TrimSpace(e.MostRelevantText)

		e.MetaDescription = strings.Replace(e.MetaDescription, r, " ", -1)
		e.MetaDescription = strings.TrimSpace(e.MetaDescription)
	}
}
