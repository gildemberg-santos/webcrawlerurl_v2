package extract

import (
	"regexp"
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
	t.normalizeHTML()
	t.Text = t.Sources.Find("body").Text()
	t.normalize()
	return t
}

func (t *Text) normalizeHTML() {
	html, _ := t.Sources.Html()
	re := regexp.MustCompile(`</(div|p|h1|h2|h3|h4|h5|h6|li|ul|ol|table|tr|td|th|blockquote|pre|code|section|article|aside|footer|header|nav|figure|figcaption|main|mark|nav|section|summary)>`)
	html = re.ReplaceAllString(html, "</$1> ")

	t.Sources, _ = goquery.NewDocumentFromReader(strings.NewReader(html))
}

func (t *Text) normalize() {
	t.Text = strings.TrimSpace(t.Text)
	t.Text = strings.Replace(t.Text, "\n", " ", -1)
	t.Text = strings.Replace(t.Text, "\t", " ", -1)
	for strings.Contains(t.Text, "  ") {
		t.Text = strings.Replace(t.Text, "  ", " ", -1)
	}
}
