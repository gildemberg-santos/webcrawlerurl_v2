package extract_test

import (
	"testing"

	"github.com/gildemberg-santos/webcrawlerurl_v2/pkg"
	"github.com/gildemberg-santos/webcrawlerurl_v2/util/extract"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestLeadsterCustom_Call(t *testing.T) {
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

	pagina := pkg.LoadPage{Url: "http://www.teste.com"}
	pagina.Load()

	leadsterCustom := extract.NewLeadsterCustom(pagina.Source, 5, 5, 30)
	response := leadsterCustom.Call()

	titule := response.TitleWebSite
	pargraph := response.MostRelevantText
	description := response.MetaDescription

	assert.Equal(t, "Titulo do site", titule)
	assert.Equal(t, "Titulo", pargraph)
	assert.Equal(t, "Meta Description", description)
}
