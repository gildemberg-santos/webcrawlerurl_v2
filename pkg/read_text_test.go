package pkg

import (
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestReadText_Call(t *testing.T) {
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

	readtext := ReadText{Url: "http://www.teste.com"}
	response, _ := readtext.Call()

	text := response.(responseSuccessReadText).ReadText.Text

	assert.Equal(t, "Titulo Paragrafo", text)
}
