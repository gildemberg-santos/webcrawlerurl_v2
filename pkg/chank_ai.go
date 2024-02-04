package pkg

import "strings"

type ChackAi struct {
	Limit      int
	ListChacks []string
	Text       string
}

func NewChackAi(limit int, text string) ChackAi {
	return ChackAi{
		Limit: limit,
		Text:  text,
	}
}

func (c *ChackAi) Call() {
	if len(c.ListChacks) < c.Limit {
		c.getListChacks()
	}
}

func (c *ChackAi) getListChacks() {
	c.ListChacks = append(strings.Split(c.Text, " "), c.ListChacks...)[0:c.Limit]
}
