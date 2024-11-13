package googleshopping_test

import (
	"testing"

	googleshopping "github.com/gildemberg-santos/webcrawlerurl_v2/util/google_shopping"
	"github.com/stretchr/testify/assert"
)

func TestRDF_NewRDF(t *testing.T) {
	t.Skip("Not implemented")
}

func TestRDF_AddItem(t *testing.T) {
	rdf := googleshopping.NewRDF()

	item1 := *googleshopping.NewItem(
		"1",
		"Product 1",
		"Description product 1",
		"Summary product 1",
		"https://www.google.com/shopping/product/1",
		"https://www.google.com/shopping/image/product/1.jpg",
		"5.00",
		"4.00",
		"in stock",
		"new",
		"male",
		"size",
		"age group",
		"color",
		*googleshopping.NewInstallment("5.00", "USD", "monthly", "12"),
	)

	item2 := *googleshopping.NewItem(
		"2",
		"Product 2",
		"Description product 2",
		"Summary product 2",
		"https://www.google.com/shopping/product/2",
		"https://www.google.com/shopping/image/product/2.jpg",
		"10.00",
		"8.00",
		"in stock",
		"new",
		"male",
		"size",
		"age group",
		"color",
		*googleshopping.NewInstallment("5.00", "USD", "monthly", "12"),
	)

	rdf.AddItem(item1)
	rdf.AddItem(item2)

	assert.Len(t, rdf.Item, 2)
	assert.Equal(t, "1", rdf.Item[0].ID.Value)
	assert.Equal(t, "2", rdf.Item[1].ID.Value)
}

func TestRDF_GetItem(t *testing.T) {
	t.Skip("Not implemented")
}
