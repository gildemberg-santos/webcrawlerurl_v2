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
