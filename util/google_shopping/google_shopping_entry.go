package googleshopping

import "fmt"

type Entry struct {
	ID                    Chardata `xml:"id"`
	Title                 Chardata `xml:"title"`
	Description           Chardata `xml:"description"`
	Link                  Chardata `xml:"link"`
	MobileLink            Chardata `xml:"mobile_link"`
	ImageLink             Chardata `xml:"image_link"`
	Price                 Chardata `xml:"price"`
	SalePrice             Chardata `xml:"sale_price"`
	Availability          Chardata `xml:"availability"`
	Condition             Chardata `xml:"condition"`
	GoogleProductCategory Chardata `xml:"google_product_category"`
}

func NewEntry(id, title, description, link, imageLink, price, salePrice, availability, condition, googleProductCategory string) *Entry {
	return &Entry{
		ID:                    *NewChardata(id),
		Title:                 *NewChardata(title),
		Description:           *NewChardata(description),
		Link:                  *NewChardata(link),
		ImageLink:             *NewChardata(imageLink),
		Price:                 *NewChardata(price),
		SalePrice:             *NewChardata(salePrice),
		Availability:          *NewChardata(availability),
		Condition:             *NewChardata(condition),
		GoogleProductCategory: *NewChardata(googleProductCategory),
	}
}

func (e *Entry) ToString() string {
	return fmt.Sprintf("ID: %s, Title: %s, Description: %s, Link: %s, ImageLink: %s, Price: %s, SalePrice: %s, Availability: %s, Condition: %s, GoogleProductCategory: %s",
		e.ID.Value, e.Title.Value, e.Description.Value, e.Link.Value, e.ImageLink.Value, e.Price.Value, e.SalePrice.Value, e.Availability.Value, e.Condition.Value, e.GoogleProductCategory.Value)
}

func (e *Entry) ToNormalise() *Entry {
	return NewEntry(e.ID.Value, e.Title.Value, e.Description.Value, e.Link.Value, e.ImageLink.Value, e.Price.Value, e.SalePrice.Value, e.Availability.Value, e.Condition.Value, e.GoogleProductCategory.Value)
}
