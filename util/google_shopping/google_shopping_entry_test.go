package googleshopping_test

import (
	"testing"

	googleshopping "github.com/gildemberg-santos/webcrawlerurl_v2/util/google_shopping"
	"github.com/stretchr/testify/assert"
)

func TestGoogleShoppingEntry_ToString(t *testing.T) {
	entry := googleshopping.NewEntry(
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

	newString := entry.ToString()

	assert.Equal(t, "ID: 1, Title: Product 1, Description: Description product 1, Link: https://www.google.com/shopping/product/1, ImageLink: https://www.google.com/shopping/image/product/1.jpg, Price: 5.00, SalePrice: 4.00, Availability: in stock, Condition: new, GoogleProductCategory: Google Product Category 1", newString)
}
