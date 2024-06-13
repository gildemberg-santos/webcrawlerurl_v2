package pkg

import (
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

// TestSmartCall_Call is a test function for the Call method of the SmartCall struct.
//
// This function tests the functionality of the Call method by mocking an HTTP GET request to "https://www.teste.com"
// and asserting the response content. It creates a new instance of the SmartCall struct with the URL "https://www.teste.com"
// and calls the Call method. It then asserts that the response is nil and that the response's Title, Paragraph, and Description
// fields match the expected values.
//
// Parameters:
// - t: The testing.T object used for running the test and reporting the results.
//
// Return type: None.
func TestSmartCall_Call(t *testing.T) {
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

	smartCall := NewSmartCall("https://www.teste.com")
	response, err := smartCall.Call()

	assert.Nil(t, err)
	assert.Equal(t, "Titulo do site", response.(responseSuccessGpt).SmartCall.Title)
	assert.Equal(t, "Titulo", response.(responseSuccessGpt).SmartCall.Paragraph)
	assert.Equal(t, "Meta Description", response.(responseSuccessGpt).SmartCall.Description)
}
