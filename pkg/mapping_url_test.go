package pkg

import (
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

// TestMappingUrl_Call is a test function for testing the Call method of the MappingUrl struct.
//
// No parameters.
// No return type.
func TestMappingUrl_Call(t *testing.T) {
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
			<a href="http://www.teste.com/teste01">Link01</h1>
			<a href="http://www.teste.com/teste02">Link02</h1>
			<a href="teste03">Link03</h1>
			<a href="/">Link04</h1>
			<a href="http://www.teste.com/teste05.pdf">Link05</h1>
			<a href="javascript:teste()">Link06</h1>
			<a href="mailto:teste@teste.com">Link07</h1>
		</body>
		<html>
	`))

	mapping_url := NewMappingUrl("http://www.teste.com", 7, true, nil)
	response, err := mapping_url.Call()

	assert.Nil(t, err)
	assert.Equal(t, true, response.Success)
	assert.Equal(t, false, response.Failure)
	assert.Equal(t, "https://www.teste.com/teste01", response.Data[0])
	assert.Equal(t, "https://www.teste.com/teste02", response.Data[1])
	assert.Equal(t, "https://www.teste.com/teste03", response.Data[2])
	assert.Equal(t, "https://www.teste.com", response.Data[3])
	assert.Len(t, response.Data, 4)
}
