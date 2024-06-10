package extract

import (
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

var doneLeadsterCustom sync.WaitGroup

type LeadsterCustom struct {
	TitleWebSite        string
	MostRelevantText    string
	MetaDescription     string
	TitleWebSiteMin     int
	MostRelevantTextMin int
	MetaDescriptionMin  int
	Source              *goquery.Document
}

func NewLeadsterCustom(source *goquery.Document, titleWebSiteMin, mostRelevantTextMin, metaDescriptionMin int) LeadsterCustom {
	return LeadsterCustom{
		Source:              source,
		TitleWebSiteMin:     titleWebSiteMin,
		MostRelevantTextMin: mostRelevantTextMin,
		MetaDescriptionMin:  metaDescriptionMin,
	}
}

func (l *LeadsterCustom) Call() *LeadsterCustom {
	doneLeadsterCustom.Add(3)
	l.extractTitleWebSite()
	l.extractMostRelevantText()
	l.extractMetaDescription()
	doneLeadsterCustom.Wait()

	l.normalize()
	return l
}

func (l *LeadsterCustom) extractTitleWebSite() {
	l.Source.Find("head").Each(func(_ int, s *goquery.Selection) {
		s.Find("title").Each(func(_ int, s *goquery.Selection) {
			text := s.Text()
			text = strings.TrimSpace(text)
			l.TitleWebSite = text
		})
	})

	if l.TitleWebSite == "" {
		l.Source.Find("title").Each(func(_ int, s *goquery.Selection) {
			text := s.Text()
			text = strings.TrimSpace(text)
			l.TitleWebSite = text
		})
	}

	l.Source.Find("meta").Each(func(_ int, s *goquery.Selection) {
		name, _ := s.Attr("property")
		if name == "title" && l.TitleWebSite == "" {
			content, _ := s.Attr("content")
			content = strings.TrimSpace(content)
			l.TitleWebSite = content
		}

		if name == "og:title" && l.TitleWebSite == "" {
			content, _ := s.Attr("content")
			content = strings.TrimSpace(content)
			l.TitleWebSite = content
		}

		if name == "twitter:title" && l.TitleWebSite == "" {
			content, _ := s.Attr("content")
			content = strings.TrimSpace(content)
			l.TitleWebSite = content
		}
	})
	defer doneLeadsterCustom.Done()
}

func (l *LeadsterCustom) extractMostRelevantText() {
	l.filterMostRelevantText("h1")

	if l.MostRelevantText != "" {
		defer doneLeadsterCustom.Done()
		return
	}

	l.filterMostRelevantText("h2")
	defer doneLeadsterCustom.Done()
}

func (l *LeadsterCustom) extractMetaDescription() {
	l.Source.Find("meta").Each(func(_ int, s *goquery.Selection) {
		name, _ := s.Attr("name")
		if name == "description" {
			content, _ := s.Attr("content")
			l.MetaDescription = content
		}

		if name == "og:description" && l.MetaDescription == "" {
			content, _ := s.Attr("content")
			l.MetaDescription = content
		}

		if name == "twitter:description" && l.MetaDescription == "" {
			content, _ := s.Attr("content")
			l.MetaDescription = content
		}

	})

	defer doneLeadsterCustom.Done()
}

func (l *LeadsterCustom) filterMostRelevantText(tag string) {
	var title = make([]string, 0)
	var first string = ""

	l.Source.Find(tag).Each(func(_ int, s *goquery.Selection) {
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
		if len(words) >= l.MostRelevantTextMin && l.MostRelevantText == "" {
			l.MostRelevantText = t
		}
	}

	if l.MostRelevantText == "" {
		l.MostRelevantText = first
	}
}

func (l *LeadsterCustom) normalize() {
	remove := []string{"\n", "\t", "\r"}

	for _, r := range remove {
		l.TitleWebSite = strings.Replace(l.TitleWebSite, r, " ", -1)
		l.TitleWebSite = strings.TrimSpace(l.TitleWebSite)

		l.MostRelevantText = strings.Replace(l.MostRelevantText, r, " ", -1)
		l.MostRelevantText = strings.TrimSpace(l.MostRelevantText)

		l.MetaDescription = strings.Replace(l.MetaDescription, r, " ", -1)
		l.MetaDescription = strings.TrimSpace(l.MetaDescription)
	}
}
