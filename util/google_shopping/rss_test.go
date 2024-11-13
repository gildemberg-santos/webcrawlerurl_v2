package googleshopping_test

import (
	"testing"

	googleshopping "github.com/gildemberg-santos/webcrawlerurl_v2/util/google_shopping"
	"github.com/stretchr/testify/assert"
)

func TestRSS_AddItem(t *testing.T) {
	rss := googleshopping.NewRSS()

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

	rss.AddItem(item1)
	rss.AddItem(item2)

	assert.Len(t, rss.Item, 2)
	assert.Equal(t, "1", rss.Item[0].ID.Value)
	assert.Equal(t, "2", rss.Item[1].ID.Value)
}

func TestRSS_GetItem(t *testing.T) {
	t.Skip("TestRSS_GetItem not implemented yet")
}
