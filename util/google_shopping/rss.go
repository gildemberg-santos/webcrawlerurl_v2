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
		if e.Description.Value != "" && e.Summary.Value == "" {
			r.Item[i].Summary = e.Description
			r.Item[i].Description = Chardata{}
		}
	}
}