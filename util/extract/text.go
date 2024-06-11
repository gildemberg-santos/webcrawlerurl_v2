package extract

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Text struct {
	Text    string
	Sources *goquery.Document
}

func NewText(source *goquery.Document) Text {
	return Text{
		Sources: source,
	}
}

func (t *Text) Call() *Text {
	t.Text = t.Sources.Find("body").Text()
	t.normalize()
	return t
}

func (t *Text) normalize() {
	t.Text = strings.TrimSpace(t.Text)
	t.Text = strings.Replace(t.Text, "\n", " ", -1)
	t.Text = strings.Replace(t.Text, "\t", " ", -1)
	for strings.Contains(t.Text, "  ") {
		t.Text = strings.Replace(t.Text, "  ", " ", -1)
	}
}
