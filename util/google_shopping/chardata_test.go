package googleshopping_test

import (
	"testing"

	googleshopping "github.com/gildemberg-santos/webcrawlerurl_v2/util/google_shopping"
	"github.com/stretchr/testify/assert"
)

func TestChardata_Normalize(t *testing.T) {
	charData := googleshopping.NewChardata("\n 1 \n")

	assert.Equal(t, "1", charData.Value)
}
