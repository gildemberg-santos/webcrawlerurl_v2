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

	LoadPage := LoadPage{Url: "http://www.teste.com"}
	err := LoadPage.Load()
	assert.Nil(t, err)
	assert.Equal(t, 200, LoadPage.StatusCode)
}

func TestLoadPage_LoadError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "http://www.teste.com", httpmock.NewStringResponder(200, ``))

	LoadPage := LoadPage{Url: "http://www.teste.com.br"}
	LoadPage.Load()
	assert.Equal(t, 404, LoadPage.StatusCode)
}
