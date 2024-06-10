package extract_test

import (
	"testing"

	"github.com/gildemberg-santos/webcrawlerurl_v2/pkg"
	"github.com/gildemberg-santos/webcrawlerurl_v2/util/extract"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestLink_Call(t *testing.T) {
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
			<a href="http://www.teste.com/teste03">Link02</h1>
		</body>
		<html>
	`))

	pagina := pkg.LoadPage{Url: "http://www.teste.com"}
	pagina.Load()

	readtext := extract.NewLink(pagina.Source, "http://www.teste.com", 2)
	response := readtext.Call()

	url01 := response.OutUrls[0]
	url02 := response.OutUrls[1]

	assert.Equal(t, "https://www.teste.com/teste01", url01)
	assert.Equal(t, "https://www.teste.com/teste02", url02)
	assert.Len(t, response.OutUrls, 2)
}
