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
		if e.Description.Value != "" && e.Summary.Value == "" {
			r.Item[i].Summary = e.Description
			r.Item[i].Description = Chardata{}
		}
	}
}
