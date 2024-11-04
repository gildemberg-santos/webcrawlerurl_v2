package googleshopping

type GooogleShoppingFeed struct {
	Entry []Entry `xml:"entry"`
}

func (g *GooogleShoppingFeed) AddEntry(e Entry) {
	g.Entry = append(g.Entry, e)
}

func (g *GooogleShoppingFeed) GetEntry() []Entry {
	return g.Entry
}
