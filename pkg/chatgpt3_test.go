package pkg

import (
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestChatGpt3_Call(t *testing.T) {
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

	LoadPage := LoadPage{Url: "http://www.teste.com"}
	LoadPage.Load()

	extractInformation := ExtractInformation{}
	extractInformation.Init(LoadPage.Source, 2, 5, 30)
	extractInformation.Call()

	assert.Equal(t, "Titulo", extractInformation.MainTitle)
	assert.Equal(t, "Paragrafo", extractInformation.MainParagraph)
	assert.Equal(t, "Meta Description", extractInformation.MetaDescription)
}
