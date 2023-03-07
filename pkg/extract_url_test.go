package pkg

import (
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestExtractUrl_Call(t *testing.T) {
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
		<a href="http://www.teste.com/teste01">Link01</h1>
		<a href="http://www.teste.com/teste02">Link02</h1>
	</body>
	<html>
	`))

	mapping_url := MappingUrl{Url: "http://www.teste.com"}
	response, _ := mapping_url.Call()

	url := response.(responseSuccessUrls).MappingUrl.Urls[0]

	assert.Equal(t, "https://www.teste.com/teste01", url)
}
