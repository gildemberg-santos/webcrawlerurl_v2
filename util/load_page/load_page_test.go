package load_page_test

import (
	"testing"

	"github.com/gildemberg-santos/webcrawlerurl_v2/util/load_page"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestLoadPage_Load(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "http://www.teste.com", httpmock.NewStringResponder(200, ``))

	loadPage := load_page.NewLoadPage("http://www.teste.com")
	err := loadPage.Call()
	assert.Nil(t, err)
	assert.Equal(t, 200, loadPage.StatusCode)
}

func TestLoadPage_LoadError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "http://www.teste.com", httpmock.NewStringResponder(404, ``))

	loadPage := load_page.NewLoadPage("http://www.teste.com")
	loadPage.Call()
	assert.Equal(t, 404, loadPage.StatusCode)
}
