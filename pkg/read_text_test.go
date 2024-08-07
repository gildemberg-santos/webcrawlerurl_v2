package pkg

import (
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

// TestReadText_Call is a test function for reading text content from a specific URL.
//
// No parameters.
// No return type.
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

	readtext := NewReadText("http://www.teste.com", nil, true)
	response, _ := readtext.Call()
	first := response.Data[0]

	assert.Equal(t, "Titulo Paragrafo", first.Text)
}
