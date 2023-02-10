package pkg

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

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
	e.extractMainTitle()
	e.extractMainParagraph()
	e.extractMetaDescription()
	e.normalize()
}

func (e *ExtractInformation) extractMainTitle() {
	e.filterTitle("h1")

	if e.MainTitle != "" {
		return
	}

	e.filterTitle("h2")
}

func (e *ExtractInformation) extractMainParagraph() {
	var paragraph = make([]string, 0)
	var first string = ""

	e.Source.Find("p").Each(func(i int, s *goquery.Selection) {
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
}

func (e *ExtractInformation) extractMetaDescription() {
	e.Source.Find("meta").Each(func(i int, s *goquery.Selection) {
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
}

func (e *ExtractInformation) filterTitle(tag string) {
	var title = make([]string, 0)
	var first string = ""

	e.Source.Find(tag).Each(func(i int, s *goquery.Selection) {
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
		if len(words) >= e.MainTitleMin && e.MainTitle == "" {
			e.MainTitle = t
		}
	}

	if e.MainTitle == "" {
		e.MainTitle = first
	}
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
