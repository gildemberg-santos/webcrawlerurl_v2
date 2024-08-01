package extract

import "github.com/PuerkitoBio/goquery"

type Metatag struct {
	Generator      string
	RetailerItemID string
	Source         *goquery.Document
}

func NewMetatag(source *goquery.Document) *Metatag {
	return &Metatag{
		Source: source,
	}
}

func (m *Metatag) Call() *Metatag {
	m.extractGenerator()
	m.extractRetailerItemID()
	return m
}

func (m *Metatag) extractGenerator() {
	m.Source.Find("meta").Each(func(_ int, s *goquery.Selection) {
		name, _ := s.Attr("name")
		if name == "generator" {
			m.Generator = s.AttrOr("content", "")
			return
		}
	})
}

func (m *Metatag) extractRetailerItemID() {
	m.Source.Find("meta").Each(func(_ int, s *goquery.Selection) {
		property, _ := s.Attr("property")
		if property == "product:retailer_item_id" {
			m.RetailerItemID = s.AttrOr("content", "")
			return
		}
	})
}
