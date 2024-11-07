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

	metaTag := NewMetaTag(t.Sources).Call()
	version := metaTag.Generator

	vtex := regexp.MustCompile(`^vtex\.render-server@8\.\d+\.\d+$`)
	chicorei := regexp.MustCompile(`^@chicorei$`)

	if vtex.MatchString(version) {
		t.Sources.Find("#desktop-top-header-id").Remove()
		t.Sources.Find("#footer-desktop-id").Remove()
		t.Sources.Find("#breadcrumb-id").Remove()
		t.Sources.Find("#tabs-product-id").Remove()
	} else if chicorei.MatchString(version) {
		t.Sources.Find("header").Remove()
		t.Sources.Find("footer").Remove()
		t.Sources.Find("#products-history-ab").Remove()
		t.Sources.Find("#advantages").Remove()
		t.Sources.Find("#creator-other-products-ab").Remove()
		t.Sources.Find("#recommendation-list").Remove()
		t.Sources.Find("#product > article > div:nth-child(9)").Remove()
		t.Sources.Find("#product > article > div:nth-child(8)").Remove()
		t.Sources.Find(".nld-chatbot").Remove()
	}
}

func (t *Text) normalize() {
	t.Text = strings.TrimSpace(t.Text)
	t.Text = strings.Replace(t.Text, "\n", " ", -1)
	t.Text = strings.Replace(t.Text, "\t", " ", -1)
	for strings.Contains(t.Text, "  ") {
		t.Text = strings.Replace(t.Text, "  ", " ", -1)
	}
}
