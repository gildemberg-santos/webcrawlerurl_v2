package extract_test

import (
	"testing"

	"github.com/gildemberg-santos/webcrawlerurl_v2/util/extract"
	"github.com/gildemberg-santos/webcrawlerurl_v2/util/load_page"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

// TestLeadsterCustom_Call is a test function for the Call method of the LeadsterCustom struct.
//
// This function tests the functionality of the Call method by mocking an HTTP GET request to "http://www.teste.com"
// and asserting the response content. It creates a new instance of the LeadsterCustom struct with the URL "http://www.teste.com"
// and calls the Call method. It then asserts that the TitleWebSite, MostRelevantText, and MetaDescription fields of the response
// match the expected values.
//
// Parameters:
// - t: The testing.T object used for running the test and reporting the results.
//
// Return type: None.
func TestLeadsterCustom_Call(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "http://www.teste.com",
		httpmock.NewStringResponder(200, `
		<!DOCTYPE html>
		<head>
			<title>Titulo do site</title>
			<meta name="description" content="Meta Description">
		</head>
		<body>
			<h1>Titulo</h1>
			<p>Paragrafo</p>
		</body>
		<html>
	`))

	page := load_page.NewLoadPage("http://www.teste.com")
	page.Call()

	leadsterCustom := extract.NewLeadsterCustom(page.Source, 5, 5, 30)
	response := leadsterCustom.Call()

	titule := response.TitleWebSite
	pargraph := response.MostRelevantText
	description := response.MetaDescription

	assert.Equal(t, "Titulo do site", titule)
	assert.Equal(t, "Titulo", pargraph)
	assert.Equal(t, "Meta Description", description)
}
