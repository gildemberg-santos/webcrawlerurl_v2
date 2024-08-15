package pkg_test

import (
	"testing"

	"github.com/gildemberg-santos/webcrawlerurl_v2/pkg"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestLeadsterAi_Call(t *testing.T) {
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

	leadsterAi := pkg.NewLeadsterAI("https://www.teste.com", "", 1, 15, true, nil)
	leadsterAi.Call(false, false)

	assert.Equal(t, "https://www.teste.com", leadsterAi.Url)
	assert.Equal(t, int64(16), leadsterAi.TotalCaracters)
	assert.Equal(t, "Titulo Paragrafo", leadsterAi.Data[0].Text)
	assert.Equal(t, int64(16), leadsterAi.Data[0].TotalCaracters)
}
