package normalize_test

import (
	"testing"

	"github.com/gildemberg-santos/webcrawlerurl_v2/util/normalize"
	"github.com/stretchr/testify/assert"
)

// TestUrl_GetUrl is a test function for testing the GetUrl method of the NewNormalizeUrl struct.
//
// No parameters.
// No return type.
func TestUrl_GetUrl(t *testing.T) {
	url, err := normalize.NewNormalizeUrl("https://www.google.com/").GetUrl()
	assert.Nil(t, err)
	assert.Equal(t, "https://www.google.com", url)
}

// TestUrl_GetUrl_Incomplete is a test function for testing the GetUrl method of the NewNormalizeUrl struct.
//
// This function tests the GetUrl method of the NewNormalizeUrl struct by passing an incomplete URL
// "www.google.com/" and asserting that the returned URL is "https://www.google.com". It also asserts that
// the error returned by the GetUrl method is nil.
//
// Parameters:
// - t: The testing.T object used for running the test and reporting the results.
//
// Return type: None.
func TestUrl_GetUrl_Incomplete(t *testing.T) {
	url, err := normalize.NewNormalizeUrl("www.google.com/").GetUrl()
	assert.Nil(t, err)
	assert.Equal(t, "https://www.google.com", url)
}

// TestUrl_GetUrl_Incomplete_Path is a test function for testing the GetUrl method of the NewNormalizeUrl struct.
//
// This function tests the GetUrl method of the NewNormalizeUrl struct by passing an incomplete URL path "teste"
// and asserting that the returned URL is "https://www.google.com/teste". It also asserts that the error returned
// by the GetUrl method is nil.
//
// Parameters:
// - t: The testing.T object used for running the test and reporting the results.
//
// Return type: None.
func TestUrl_GetUrl_Incomplete_Path(t *testing.T) {
	uri := normalize.NewNormalizeUrl("teste")
	uri.BaseUrl = "https://www.google.com"
	url, err := uri.GetUrl()

	assert.Nil(t, err)
	assert.Equal(t, "https://www.google.com/teste", url)
}

func TestUrl_GetUrl_SiteMap(t *testing.T) {
	url, err := normalize.NewNormalizeUrl("https://www.google.com/sitemap.xml").GetUrl()
	assert.Nil(t, err)
	assert.Equal(t, "https://www.google.com/sitemap.xml", url)
}

func TestUrl_MD5(t *testing.T) {
	hash := normalize.NewNormalizeUrl("https://www.google.com/contato/?p=google").MD5()
	assert.Equal(t, "9635180a42da30ac7cd3f4b0447975b9", hash)
}
