package pkg

import (
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestExtractText_Call(t *testing.T) {
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

	pagina := LoadPage{Url: "http://www.teste.com"}
	pagina.Load()

	readtext := ExtractText{}
	readtext.Init(pagina.Source)
	response, _ := readtext.Call()

	text := response.(ExtractText).Text

	assert.Equal(t, "Titulo Paragrafo", text)
}
