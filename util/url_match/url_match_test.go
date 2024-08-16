package url_match_test

import (
	"testing"

	"github.com/gildemberg-santos/webcrawlerurl_v2/util/url_match"
	"github.com/stretchr/testify/assert"
)

func TestUrlMatch_Call(t *testing.T) {
	assert.Equal(t, true, url_match.NewUrlMatch("https://example.com/perfil/**").Call("https://example.com/perfil/contact"))
	assert.Equal(t, false, url_match.NewUrlMatch("https://example.com/perfil/**").Call("https://example.com/admin/contact"))
	assert.Equal(t, true, url_match.NewUrlMatch("https://**/perfil/**").Call("https://example.com/perfil/contact"))
	assert.Equal(t, false, url_match.NewUrlMatch("https://**/perfil/**").Call("https://example.com/admin/contact"))
	assert.Equal(t, true, url_match.NewUrlMatch("https://example.com/**/contact").Call("https://example.com/perfil/contact"))
	assert.Equal(t, false, url_match.NewUrlMatch("https://example.com/**/contact").Call("https://example.com/perfil/admin"))
}
