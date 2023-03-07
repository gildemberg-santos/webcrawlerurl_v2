package pkg

import (
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestExtractInformation_Call(t *testing.T) {
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

	chatgpt := ChatGpt3{Url: "http://www.teste.com"}
	response, _ := chatgpt.Call()

	titule := response.(responseSuccessGpt).ChatGpt.Title
	pargraph := response.(responseSuccessGpt).ChatGpt.Paragraph
	description := response.(responseSuccessGpt).ChatGpt.Description

	assert.Equal(t, "Titulo", titule)
	assert.Equal(t, "Paragrafo", pargraph)
	assert.Equal(t, "Meta Description", description)
}
