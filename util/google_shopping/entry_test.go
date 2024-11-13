package googleshopping_test

import (
	"testing"

	googleshopping "github.com/gildemberg-santos/webcrawlerurl_v2/util/google_shopping"
	"github.com/stretchr/testify/assert"
)

func TestEntry_NewEntry(t *testing.T) {
	entry := googleshopping.NewEntry(
		"123", "Test Title", "Test Description", "Test Summary", "http://example.com",
		"http://example.com/image.jpg", "100", "80", "in stock", "new", "unisex", "L", "adult", "red",
		googleshopping.Installment{Months: googleshopping.Chardata{Value: "12"}, Amount: googleshopping.Chardata{Value: "10"}, Downpayment: googleshopping.Chardata{Value: "0"}, CreditType: googleshopping.Chardata{Value: "credit"}},
	)

	assert.Equal(t, "123", entry.ID.Value)
	assert.Equal(t, "Test Title", entry.Title.Value)
	assert.Equal(t, "Test Description", entry.Description.Value)
	assert.Equal(t, "Test Summary", entry.Summary.Value)
	assert.Equal(t, "http://example.com", entry.Link.Value)
	assert.Equal(t, "http://example.com/image.jpg", entry.ImageLink.Value)
	assert.Equal(t, "100", entry.Price.Value)
	assert.Equal(t, "80", entry.SalePrice.Value)
	assert.Equal(t, "in stock", entry.Availability.Value)
	assert.Equal(t, "new", entry.Condition.Value)
	assert.Equal(t, "unisex", entry.Gender.Value)
	assert.Equal(t, "L", entry.Size.Value)
	assert.Equal(t, "adult", entry.AgeGroup.Value)
	assert.Equal(t, "red", entry.Color.Value)
	assert.Equal(t, "12", entry.Installment.Months.Value)
	assert.Equal(t, "10", entry.Installment.Amount.Value)
	assert.Equal(t, "0", entry.Installment.Downpayment.Value)
	assert.Equal(t, "credit", entry.Installment.CreditType.Value)
}

func TestEntry_ToNormalise(t *testing.T) {
	entry := googleshopping.NewEntry(
		"123", "Test Title", "Test Description", "Test Summary", "http://example.com",
		"http://example.com/image.jpg", "100", "80", "in stock", "new", "unisex", "L", "adult", "red",
		googleshopping.Installment{Months: googleshopping.Chardata{Value: "12"}, Amount: googleshopping.Chardata{Value: "10"}, Downpayment: googleshopping.Chardata{Value: "0"}, CreditType: googleshopping.Chardata{Value: "credit"}},
	)

	normalisedEntry := entry.ToNormalise()

	assert.Equal(t, entry.ID.Value, normalisedEntry.ID.Value)
	assert.Equal(t, entry.Title.Value, normalisedEntry.Title.Value)
	assert.Equal(t, entry.Description.Value, normalisedEntry.Description.Value)
	assert.Equal(t, entry.Summary.Value, normalisedEntry.Summary.Value)
	assert.Equal(t, entry.Link.Value, normalisedEntry.Link.Value)
	assert.Equal(t, entry.ImageLink.Value, normalisedEntry.ImageLink.Value)
	assert.Equal(t, entry.Price.Value, normalisedEntry.Price.Value)
	assert.Equal(t, entry.SalePrice.Value, normalisedEntry.SalePrice.Value)
	assert.Equal(t, entry.Availability.Value, normalisedEntry.Availability.Value)
	assert.Equal(t, entry.Condition.Value, normalisedEntry.Condition.Value)
	assert.Equal(t, entry.Gender.Value, normalisedEntry.Gender.Value)
	assert.Equal(t, entry.Size.Value, normalisedEntry.Size.Value)
	assert.Equal(t, entry.AgeGroup.Value, normalisedEntry.AgeGroup.Value)
	assert.Equal(t, entry.Color.Value, normalisedEntry.Color.Value)
	assert.Equal(t, entry.Installment.Months.Value, normalisedEntry.Installment.Months.Value)
	assert.Equal(t, entry.Installment.Amount.Value, normalisedEntry.Installment.Amount.Value)
	assert.Equal(t, entry.Installment.Downpayment.Value, normalisedEntry.Installment.Downpayment.Value)
	assert.Equal(t, entry.Installment.CreditType.Value, normalisedEntry.Installment.CreditType.Value)
}
