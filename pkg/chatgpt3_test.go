package pkg

import (
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestChatGpt3_Call(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://www.teste.com",
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

	chatgpt3 := ChatGpt3{Url: "https://www.teste.com"}
	response, err := chatgpt3.Call()

	assert.Nil(t, err)
	assert.Equal(t, "Titulo", response.(responseSuccessGpt).ChatGpt.Title)
	assert.Equal(t, "Paragrafo", response.(responseSuccessGpt).ChatGpt.Paragraph)
	assert.Equal(t, "Meta Description", response.(responseSuccessGpt).ChatGpt.Description)
}
