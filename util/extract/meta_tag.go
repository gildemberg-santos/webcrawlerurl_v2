package extract

import "github.com/PuerkitoBio/goquery"

type MetaTag struct {
	Generator      string            `json:"generator"`
	RetailerItemID string            `json:"retailer_item_id"`
	Source         *goquery.Document `json:"-"`
}

func NewMetaTag(source *goquery.Document) *MetaTag {
	return &MetaTag{
		Source: source,
	}
}

func (m *MetaTag) Call() *MetaTag {
	m.extractGenerator()
	m.extractRetailerItemID()
	return m
}

func (m *MetaTag) extractGenerator() {
	m.Source.Find("meta").Each(func(_ int, s *goquery.Selection) {
		name, _ := s.Attr("name")
		if name == "generator" {
			m.Generator = s.AttrOr("content", "")
			return
		} else if name == "twitter:creator" {
			m.Generator = s.AttrOr("content", "")
			return
		}
	})
}

func (m *MetaTag) extractRetailerItemID() {
	m.Source.Find("meta").Each(func(_ int, s *goquery.Selection) {
		property, _ := s.Attr("property")
		if property == "product:retailer_item_id" {
			m.RetailerItemID = s.AttrOr("content", "")
			return
		}
	})
}
