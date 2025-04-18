package googleshopping

type RDF struct {
	Item []Item `xml:"item"`
}

func NewRDF() *RDF {
	return &RDF{}
}

func (r *RDF) AddItem(i Item) {
	r.Item = append(r.Item, i)
}

func (r *RDF) GetItem() []Item {
	return r.Item
}

func (r *RDF) Normalize() {
	for i, e := range r.Item {
		if e.Summary.Value != "" && e.Description.Value == "" {
			r.Item[i].Description = e.Summary
			r.Item[i].Summary = Chardata{}
		}
	}
}
