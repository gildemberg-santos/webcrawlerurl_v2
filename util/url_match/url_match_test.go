package url_match_test

import (
	"testing"

	"github.com/gildemberg-santos/webcrawlerurl_v2/util/url_match"
	"github.com/stretchr/testify/assert"
)

func TestUrlMatch_Call(t *testing.T) {
	filterValid := url_match.NewUrlMatch("https://example.com/perfil/**")
	filterInvalid := url_match.NewUrlMatch("https://example.com/perfil/**")
	assert.Equal(t, filterValid.Call("https://example.com/perfil/contact"), true)
	assert.Equal(t, filterInvalid.Call("https://example.com/admin/contact"), false)
}
