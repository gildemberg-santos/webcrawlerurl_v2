package googleshopping_test

import (
	"testing"

	googleshopping "github.com/gildemberg-santos/webcrawlerurl_v2/util/google_shopping"
	"github.com/stretchr/testify/assert"
)

func TestGooogleShoppingFeed_AddEntry(t *testing.T) {
	feed := googleshopping.NewFeed()

	entry1 := *googleshopping.NewEntry(
		"1",
		"Product 1",
		"Description product 1",
		"https://www.google.com/shopping/product/1",
		"https://www.google.com/shopping/image/product/1.jpg",
		"5.00",
		"4.00",
		"in stock",
		"new",
		"Google Product Category 1",
	)

	entry2 := *googleshopping.NewEntry(
		"2",
		"Product 2",
		"Description product 2",
		"https://www.google.com/shopping/product/2",
		"https://www.google.com/shopping/image/product/2.jpg",
		"10.00",
		"8.00",
		"in stock",
		"new",
		"Google Product Category 2",
	)

	feed.AddEntry(entry1)
	feed.AddEntry(entry2)

	assert.Len(t, feed.Entry, 2)
	assert.Equal(t, "1", feed.Entry[0].ID.Value)
	assert.Equal(t, "2", feed.Entry[1].ID.Value)
}

func TestGooogleShoppingFeed_GetEntry(t *testing.T) {
	feed := googleshopping.NewFeed()

	entry1 := *googleshopping.NewEntry(
		"1",
		"Product 1",
		"Description product 1",
		"https://www.google.com/shopping/product/1",
		"https://www.google.com/shopping/image/product/1.jpg",
		"5.00",
		"4.00",
		"in stock",
		"new",
		"Google Product Category 1",
	)

	entry2 := *googleshopping.NewEntry(
		"2",
		"Product 2",
		"Description product 2",
		"https://www.google.com/shopping/product/2",
		"https://www.google.com/shopping/image/product/2.jpg",
		"10.00",
		"8.00",
		"in stock",
		"new",
		"Google Product Category 2",
	)

	feed.AddEntry(entry1)
	feed.AddEntry(entry2)

	entries := feed.GetEntry()

	assert.Len(t, entries, 2)
	assert.Equal(t, "1", entries[0].ID.Value)
	assert.Equal(t, "2", entries[1].ID.Value)
}
