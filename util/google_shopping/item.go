package googleshopping

type Item struct {
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

func NewItem(id, title, description, summary, link, imageLink, price, salePrice, availability, condition, gender, size, ageGroup, color string, installment Installment) *Item {
	return &Item{
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

func (i *Item) ToNormalise() *Item {
	return NewItem(i.ID.Value, i.Title.Value, i.Description.Value, i.Summary.Value, i.Link.Value, i.ImageLink.Value, i.Price.Value, i.SalePrice.Value, i.Availability.Value, i.Condition.Value, i.Gender.Value, i.Size.Value, i.AgeGroup.Value, i.Color.Value, i.Installment)
}
