package chunck

import "strings"

type Chunck struct {
	Text        string
	Limit       int64
	ListChuncks []string
	CountChunck int64
}

func NewChunck(text string, limit int64) Chunck {
	return Chunck{
		Text:  text,
		Limit: limit,
	}
}

func (c *Chunck) Call() *Chunck {
	if int64(len(c.ListChuncks)) < c.Limit {
		c.getListChuncks()
		c.removeDuplicateChuncks()
	}
	c.CountChunck = int64(len(c.ListChuncks))
	return c
}

func (c *Chunck) getListChuncks() {
	text := c.Text
	for len(text) > 0 {
		if int64(len(c.ListChuncks)) == c.Limit {
			break
		}
		if len(text) <= 100 {
			c.ListChuncks = append(c.ListChuncks, text)
			break
		}
		i := strings.LastIndex(text[:100], " ")
		if i == -1 && len(text) >= 100 {
			i = 100
		}
		c.ListChuncks = append(c.ListChuncks, text[:i])
		text = text[i+1:]
	}
}

func (c *Chunck) removeDuplicateChuncks() {
	var listChuncks []string
	for _, chunck := range c.ListChuncks {
		if !contains(listChuncks, chunck) {
			listChuncks = append(listChuncks, chunck)
		}
	}
	c.ListChuncks = listChuncks
}

func contains(list []string, chunck string) bool {
	for _, item := range list {
		if item == chunck {
			return true
		}
	}
	return false
}
