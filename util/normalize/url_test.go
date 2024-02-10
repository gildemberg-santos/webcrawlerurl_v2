package normalize_test

import (
	"testing"

	"github.com/gildemberg-santos/webcrawlerurl_v2/util/normalize"
	"github.com/stretchr/testify/assert"
)

func TestUrl_GetUrl(t *testing.T) {
	url, err := normalize.NewNormalizeUrl("www.google.com").GetUrl()
	assert.Nil(t, err)
	assert.Equal(t, "https://www.google.com", url)
}

func TestUrl_GetUrl_Incomplete(t *testing.T) {
	url, err := normalize.NewNormalizeUrl("www.google.com").GetUrl()
	assert.Nil(t, err)
	assert.Equal(t, "https://www.google.com", url)
}
