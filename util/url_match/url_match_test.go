package url_match_test

import (
	"testing"

	"github.com/gildemberg-santos/webcrawlerurl_v2/util/url_match"
	"github.com/stretchr/testify/assert"
)

func TestUrlMatch_Call(t *testing.T) {
	assert.Equal(t, url_match.NewUrlMatch("https://example.com/perfil/**").Call("https://example.com/perfil/contact"), true)
	assert.Equal(t, url_match.NewUrlMatch("https://example.com/perfil/**").Call("https://example.com/admin/contact"), false)
	assert.Equal(t, url_match.NewUrlMatch("https://example.com**").Call("https://example.com"), true)
	assert.Equal(t, url_match.NewUrlMatch("https://example.com/**").Call("https://example.com"), false)
}
