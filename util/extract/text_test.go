package extract_test

import (
	"testing"

	"github.com/gildemberg-santos/webcrawlerurl_v2/util/extract"
	"github.com/gildemberg-santos/webcrawlerurl_v2/util/load_page"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestText_Call(t *testing.T) {
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

	readtext := extract.NewText(page.Source)
	response := readtext.Call()

	text := response.Text

	assert.Equal(t, "Titulo Paragrafo", text)
}
