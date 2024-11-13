package googleshopping_test

import (
	"testing"

	googleshopping "github.com/gildemberg-santos/webcrawlerurl_v2/util/google_shopping"
	"github.com/stretchr/testify/assert"
)

func TestChardata_NewChardata(t *testing.T) {
	charData := googleshopping.NewChardata("1")

	assert.New(t).Equal("1", charData.Value)
}

func TestChardata_Normalize(t *testing.T) {
	charData := googleshopping.NewChardata("\n 1 \n")

	assert.Equal(t, "1", charData.Value)
}

func TestChardata_EmptyString(t *testing.T) {
	charData := googleshopping.NewChardata("")

	assert.Equal(t, "", charData.Value)
}

func TestChardata_OnlySpaces(t *testing.T) {
	charData := googleshopping.NewChardata("   ")

	assert.Equal(t, "", charData.Value)
}

func TestChardata_OnlyNewlines(t *testing.T) {
	charData := googleshopping.NewChardata("\n\n\n")

	assert.Equal(t, "", charData.Value)
}

func TestChardata_SpacesAndNewlines(t *testing.T) {
	charData := googleshopping.NewChardata(" \n  \n ")

	assert.Equal(t, "", charData.Value)
}

func TestChardata_MixedContent(t *testing.T) {
	charData := googleshopping.NewChardata(" \n 1 \n 2 \n ")

	assert.Equal(t, "1 2", charData.Value)
}
