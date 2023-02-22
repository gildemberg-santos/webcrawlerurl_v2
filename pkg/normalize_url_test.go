package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalizeUrl_GetUrl(t *testing.T) {
	normalizeUrl := NormalizeUrl{Url: "www.google.com"}
	url, err := normalizeUrl.GetUrl()
	assert.Nil(t, err)
	assert.Equal(t, "https://www.google.com", url)
}
