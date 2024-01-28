package pkg

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type ExtractText struct {
	Text   string
	Source *goquery.Document
}

func (e *ExtractText) Init(source *goquery.Document) {
	e.Source = source
}

func (e ExtractText) Call() (interface{}, error) {
	e.Text = e.Source.Find("body").Text()
	e.normalize()
	return e, nil
}

func (e *ExtractText) normalize() {
	e.Text = strings.TrimSpace(e.Text)
	e.Text = strings.Replace(e.Text, "\n", " ", -1)
	e.Text = strings.Replace(e.Text, "\t", " ", -1)
	e.Text = strings.Replace(e.Text, "   ", " ", -1)
	e.Text = strings.Replace(e.Text, "  ", " ", -1)
}
