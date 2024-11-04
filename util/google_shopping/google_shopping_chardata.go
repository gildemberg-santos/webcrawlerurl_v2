package googleshopping

import "strings"

type Chardata struct {
	Value string `xml:",chardata"`
}

func NewChardata(value string) *Chardata {
	var c Chardata
	c.Value = value
	c.Normalize()
	return &c
}

func (c *Chardata) Normalize() {
	c.Value = strings.TrimSpace(strings.ReplaceAll(c.Value, "\n", ""))
}
