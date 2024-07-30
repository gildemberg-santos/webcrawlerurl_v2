package load_page_test

import (
	"testing"

	"github.com/gildemberg-santos/webcrawlerurl_v2/util/load_page"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

// TestLoadPage_Load is a test function for loading a page.
//
// Parameters:
// - t: The testing.T object used for running the test and reporting the results.
// Return type: None.
func TestLoadPage_Load(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "http://www.teste.com", httpmock.NewStringResponder(200, ``))

	loadPage := load_page.NewLoadPage("http://www.teste.com", true)
	err := loadPage.Call()
	assert.Nil(t, err)
	assert.Equal(t, 200, loadPage.StatusCode)
}

// TestLoadPage_LoadError is a test function for testing the behavior of the LoadPage struct's Call method when the HTTP request returns a 404 status code.
//
// It activates the HTTP mock, registers a responder for the specified URL with a 404 status code, creates a new LoadPage instance with the URL, calls the Call method, and asserts that the StatusCode is equal to 404.
//
// Parameters:
// - t: The testing.T object used for running the test and reporting the results.
//
// Return type: None.
func TestLoadPageFast_LoadError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "http://www.teste.com", httpmock.NewStringResponder(404, ``))

	loadPage := load_page.NewLoadPage("http://www.teste.com", true)
	loadPage.Call()
	assert.Equal(t, 404, loadPage.StatusCode)
}

// func TestLoadPageSlow_Load(t *testing.T) {
// 	httpmock.Activate()
// 	defer httpmock.DeactivateAndReset()

// 	httpmock.RegisterResponder("GET", "http://www.teste.com", httpmock.NewStringResponder(200, ``))

// 	loadPage := load_page.NewLoadPage("http://www.teste.com", false)
// 	err := loadPage.Call()
// 	assert.Nil(t, err)
// 	assert.Equal(t, 200, loadPage.StatusCode)
// }
