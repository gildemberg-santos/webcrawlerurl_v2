package googleshopping

type Feed struct {
	Entry []Entry `xml:"entry"`
}

func NewFeed() *Feed {
	return &Feed{}
}

func (f *Feed) AddEntry(e Entry) {
	f.Entry = append(f.Entry, e)
}

func (f *Feed) GetEntry() []Entry {
	return f.Entry
}

func (f *Feed) Normalize() {
	for i, e := range f.Entry {
		if e.Summary.Value != "" && e.Description.Value == "" {
			f.Entry[i].Description = e.Summary
			f.Entry[i].Summary = Chardata{}
		}
	}
}
