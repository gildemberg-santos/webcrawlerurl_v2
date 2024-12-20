package googleshopping

type Entry struct {
	ID           Chardata    `xml:"id"`
	Title        Chardata    `xml:"title"`
	Description  Chardata    `xml:"description"`
	Summary      Chardata    `xml:"summary"`
	Link         Chardata    `xml:"link"`
	MobileLink   Chardata    `xml:"mobile_link"`
	ImageLink    Chardata    `xml:"image_link"`
	Price        Chardata    `xml:"price"`
	SalePrice    Chardata    `xml:"sale_price"`
	Availability Chardata    `xml:"availability"`
	Condition    Chardata    `xml:"condition"`
	Gender       Chardata    `xml:"gender"`
	Size         Chardata    `xml:"size"`
	AgeGroup     Chardata    `xml:"age_group"`
	Color        Chardata    `xml:"color"`
	Installment  Installment `xml:"installment"`
}

func NewEntry(id, title, description, summary, link, imageLink, price, salePrice, availability, condition, gender, size, ageGroup, color string, installment Installment) *Entry {
	return &Entry{
		ID:           *NewChardata(id),
		Title:        *NewChardata(title),
		Description:  *NewChardata(description),
		Summary:      *NewChardata(summary),
		Link:         *NewChardata(link),
		ImageLink:    *NewChardata(imageLink),
		Price:        *NewChardata(price),
		SalePrice:    *NewChardata(salePrice),
		Availability: *NewChardata(availability),
		Condition:    *NewChardata(condition),
		Gender:       *NewChardata(gender),
		Size:         *NewChardata(size),
		AgeGroup:     *NewChardata(ageGroup),
		Color:        *NewChardata(color),
		Installment:  *NewInstallment(installment.Months.Value, installment.Amount.Value, installment.Downpayment.Value, installment.CreditType.Value),
	}
}

func (e *Entry) ToNormalise() *Entry {
	return NewEntry(e.ID.Value, e.Title.Value, e.Description.Value, e.Summary.Value, e.Link.Value, e.ImageLink.Value, e.Price.Value, e.SalePrice.Value, e.Availability.Value, e.Condition.Value, e.Gender.Value, e.Size.Value, e.AgeGroup.Value, e.Color.Value, e.Installment)
}
