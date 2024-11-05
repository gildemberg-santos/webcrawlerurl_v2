package googleshopping

type Entry struct {
	ID           Chardata `xml:"id"`
	Title        Chardata `xml:"title"`
	Description  Chardata `xml:"description"`
	Link         Chardata `xml:"link"`
	ImageLink    Chardata `xml:"image_link"`
	Price        Chardata `xml:"price"`
	Availability Chardata `xml:"availability"`
}

func NewEntry(id, title, description, link, imageLink, price, availability string) *Entry {
	return &Entry{
		ID:           *NewChardata(id),
		Title:        *NewChardata(title),
		Description:  *NewChardata(description),
		Link:         *NewChardata(link),
		ImageLink:    *NewChardata(imageLink),
		Price:        *NewChardata(price),
		Availability: *NewChardata(availability),
	}
}

func (e *Entry) ToString() string {
	return "ID: " + e.ID.Value + ", Title: " + e.Title.Value + ", Description: " + e.Description.Value + ", Link: " + e.Link.Value + ", ImageLink: " + e.ImageLink.Value + ", Price: " + e.Price.Value + ", Availability: " + e.Availability.Value
}

func (e *Entry) ToNormalise() *Entry {
	return NewEntry(e.ID.Value, e.Title.Value, e.Description.Value, e.Link.Value, e.ImageLink.Value, e.Price.Value, e.Availability.Value)
}
