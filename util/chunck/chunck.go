package chunck

import "strings"

type Chunck struct {
	Limit       int
	ListChuncks []string
	Text        string
}

func NewChunck(limit int, text string) Chunck {
	return Chunck{
		Limit: limit,
		Text:  text,
	}
}

func (c *Chunck) Call() *Chunck {
	if len(c.ListChuncks) < c.Limit {
		c.getListChuncks()
	}
	return c
}

func (c *Chunck) getListChuncks() {
	c.ListChuncks = append(strings.Split(c.Text, " "), c.ListChuncks...)[0:c.Limit]
}
