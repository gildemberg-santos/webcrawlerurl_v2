package googleshopping

type RSS struct {
	Item []Item `xml:"item"`
}

func NewRSS() *RSS {
	return &RSS{}
}

func (r *RSS) AddItem(i Item) {
	r.Item = append(r.Item, i)
}

func (r *RSS) GetItem() []Item {
	return r.Item
}

func (r *RSS) Normalize() {
	for i, e := range r.Item {
		if e.Summary.Value != "" && e.Description.Value == "" {
			r.Item[i].Description = e.Summary
			r.Item[i].Summary = Chardata{}
		}
	}
}
