package extract_test

import (
	"testing"

	"github.com/gildemberg-santos/webcrawlerurl_v2/util/extract"
	"github.com/gildemberg-santos/webcrawlerurl_v2/util/load_page"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestMetatag_Call(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "http://www.exemple.com",
		httpmock.NewStringResponder(200, `
		<!DOCTYPE html>
		<head>
			<title>Titulo do site</title>
			<meta name="generator" content="vtex.render-server@8.172.2">
			<meta property="product:retailer_item_id" content="AG2708010" data-react-helmet="true">
			<meta name="description" content="Meta Description">
		</head>
		<body>
			<h1>Titulo</h1>
			<p>Paragrafo</p>
		</body>
		<html>
	`))

	page := load_page.NewLoadPage("http://www.exemple.com", true)
	page.Call()

	metatag := extract.NewMetaTag(page.Source)
	response := metatag.Call()
	assert.Equal(t, "vtex.render-server@8.172.2", response.Generator)
	assert.Equal(t, "AG2708010", response.RetailerItemID)
}
