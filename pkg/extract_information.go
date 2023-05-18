package pkg

import (
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

var doneExtractInformation sync.WaitGroup

type ExtractInformation struct {
	MainTitle          string
	MainParagraph      string
	MetaDescription    string
	MainTitleMin       int
	MainParagraphMin   int
	MetaDescriptionMin int
	Source             *goquery.Document
}

func (e *ExtractInformation) Init(source *goquery.Document, titleMin, paragraphMin, descriptionMin int) {
	e.Source = source
	e.MainTitleMin = titleMin
	e.MainParagraphMin = paragraphMin
	e.MetaDescriptionMin = descriptionMin
}

func (e *ExtractInformation) Call() {
	doneExtractInformation.Add(3)
	go e.extractMainTitle()
	go e.extractMainParagraph()
	go e.extractMetaDescription()
	doneExtractInformation.Wait()

	e.normalize()
}

func (e *ExtractInformation) extractMainTitle() {
	e.Source.Find("title").Each(func(_ int, s *goquery.Selection) {
		text := s.Text()
		text = strings.TrimSpace(text)
		e.MainTitle = text
	})
	defer doneExtractInformation.Done()
}

func (e *ExtractInformation) extractMainParagraph() {
	var paragraph = make([]string, 0)
	var first string = ""

	e.Source.Find("p").Each(func(_ int, s *goquery.Selection) {
		text := s.Text()

		text = strings.TrimSpace(text)

		if text != "" {
			if first == "" {
				first = text
			}
			paragraph = append(paragraph, text)
		}
	})

	for _, p := range paragraph {
		words := strings.Split(p, " ")
		if len(words) >= e.MainParagraphMin {
			e.MainParagraph = p
			break
		}
	}

	if e.MainParagraph == "" {
		e.MainParagraph = first
	}

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


func (e *ExtractInformation) normalize() {
	remove := []string{"\n", "\t", "\r"}

	for _, r := range remove {
		e.MainTitle = strings.Replace(e.MainTitle, r, " ", -1)
		e.MainTitle = strings.TrimSpace(e.MainTitle)

		e.MainParagraph = strings.Replace(e.MainParagraph, r, " ", -1)
		e.MainParagraph = strings.TrimSpace(e.MainParagraph)

		e.MetaDescription = strings.Replace(e.MetaDescription, r, " ", -1)
		e.MetaDescription = strings.TrimSpace(e.MetaDescription)
	}
}
