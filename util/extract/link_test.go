package extract_test

import (
	"testing"

	"github.com/gildemberg-santos/webcrawlerurl_v2/util/extract"
	"github.com/gildemberg-santos/webcrawlerurl_v2/util/load_page"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

// TestLink_Call is a test function for testing the Call method of the Link struct.
//
// It activates HTTP mock, registers responders for specific URLs, creates a new LoadPage, calls it, creates a new Link,
// calls it, checks the extracted URLs, and asserts the expected values.
func TestLink_Call(t *testing.T) {
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
			<a href="http://www.teste.com/teste01">Link01</h1>
			<a href="http://www.teste.com/teste02">Link02</h1>
			<a href="teste03">Link03</h1>
			<a href="http://www.teste.com/teste04">Link04</h1>
		</body>
		<html>
	`))

	page := load_page.NewLoadPage("http://www.teste.com")
	page.Call()

	readtext := extract.NewLink(page.Source, "http://www.teste.com", 3)
	response := readtext.Call()

	url01 := response.OutUrls[0]
	url02 := response.OutUrls[1]
	url03 := response.OutUrls[2]

	assert.Equal(t, "https://www.teste.com/teste01", url01)
	assert.Equal(t, "https://www.teste.com/teste02", url02)
	assert.Equal(t, "https://www.teste.com/teste03", url03)
	assert.Len(t, response.OutUrls, 3)
}
