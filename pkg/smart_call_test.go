package pkg

import (
	"errors"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"testing"
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

	smartCall := NewSmartCall("https://www.teste.com", true)
	response, err := smartCall.Call()

	assert.Nil(t, err)
	assert.Equal(t, "Titulo do site", response.(responseSuccessGpt).SmartCall.Title)
	assert.Equal(t, "Titulo", response.(responseSuccessGpt).SmartCall.Paragraph)
	assert.Equal(t, "Meta Description", response.(responseSuccessGpt).SmartCall.Description)
}

// TestSmartCall_Call_EmptyUrl tests the Call method of the SmartCall struct when the URL is empty.
// It checks that the method returns an error and the appropriate error message.
func TestSmartCall_Call_EmptyUrl(t *testing.T) {
	smartCall := NewSmartCall("", true)
	response, err := smartCall.Call()

	assert.NotNil(t, err)
	assert.Equal(t, "url is empty", err.Error())
	assert.Equal(t, "", response.(responseErroGpt).Url)
	assert.Equal(t, 500, response.(responseErroGpt).StatusCode)
}

// TestSmartCall_Call_InvalidUrl tests the Call method of the SmartCall struct when the URL is invalid.
// It checks that the method returns an error and the appropriate error message.
func TestSmartCall_Call_InvalidUrl(t *testing.T) {
	smartCall := NewSmartCall("invalid-url", true)
	response, err := smartCall.Call()

	assert.NotNil(t, err)
	assert.Equal(t, "url is invalid", err.Error())
	assert.Equal(t, "invalid-url", response.(responseErroGpt).Url)
	assert.Equal(t, 500, response.(responseErroGpt).StatusCode)
}

// TestSmartCall_Call_ValidUrl tests the Call method of the SmartCall struct with a valid URL.
func TestSmartCall_Call_ValidUrl(t *testing.T) {
	smartCall := NewSmartCall("https://www.valid-url.com", true)
	response, err := smartCall.Call()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, "https://www.valid-url.com", response.(responseSuccessGpt).Url)
	assert.Equal(t, 200, response.(responseSuccessGpt).StatusCode)
}

// TestSmartCall_Call_LoadPageFast tests the Call method of the SmartCall struct with LoadPageFast set to true.
func TestSmartCall_Call_LoadPageFast(t *testing.T) {
	smartCall := NewSmartCall("https://www.fast-url.com", true)
	response, err := smartCall.Call()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, "https://www.fast-url.com", response.(responseSuccessGpt).Url)
	assert.Equal(t, 200, response.(responseSuccessGpt).StatusCode)
	assert.True(t, smartCall.LoadPageFast) // Ensure LoadPageFast is true
}

// TestSmartCall_Call_Timeout tests the Call method of the SmartCall struct with a timeout scenario.
func TestSmartCall_Call_Timeout(t *testing.T) {
	httpmock.Activate()

	defer httpmock.DeactivateAndReset()
	httpmock.NewErrorResponder(errors.New("Client.Timeout exceeded while awaiting headers"))

	smartCall := NewSmartCall("https://www.timeout-url.com", true)
	response, err := smartCall.Call()

	assert.NotNil(t, err)
	assert.Equal(t, "Error to send request -> Get \"https://www.timeout-url.com\": no responder found", err.Error())
	assert.Equal(t, "https://www.timeout-url.com", response.(responseErroGpt).Url)
	assert.Equal(t, 404, response.(responseErroGpt).StatusCode)
	assert.Nil(t, response.(responseErroGpt).SmartCall)
}

// TestSmartCall_Call_Error tests the Call method of the SmartCall struct when an error occurs during the call.
func TestSmartCall_Call_Error(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://www.error-url.com",
		httpmock.NewErrorResponder(errors.New("Network error")))

	smartCall := NewSmartCall("https://www.error-url.com", true)
	response, err := smartCall.Call()

	assert.NotNil(t, err)
	assert.Equal(t, "Error to send request -> Get \"https://www.error-url.com\": Network error", err.Error())
	assert.Equal(t, "https://www.error-url.com", response.(responseErroGpt).Url)
	assert.Equal(t, 404, response.(responseErroGpt).StatusCode)
	assert.Nil(t, response.(responseErroGpt).SmartCall)
}
