package pkg

import (
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestLoadPage_Load(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "http://www.teste.com", httpmock.NewStringResponder(200, ``))

	loadPage := NewLoadPage("http://www.teste.com")
	err := loadPage.Call()
	assert.Nil(t, err)
	assert.Equal(t, 200, loadPage.StatusCode)
}

func TestLoadPage_LoadError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "http://www.teste.com", httpmock.NewStringResponder(404, ``))

	loadPage := NewLoadPage("http://www.teste.com")
	loadPage.Call()
	assert.Equal(t, 404, loadPage.StatusCode)
}
