package pkg_test

import (
	"testing"

	"github.com/gildemberg-santos/webcrawlerurl_v2/pkg"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestEcommerce_Call(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://www.teste.com/ah0116001-sandalia-basica-couro-preta-detalhe-laser/p", httpmock.NewStringResponder(200, `
		<!DOCTYPE html>
			<head>
				<title>Titulo do site</title>
				<meta property="product:retailer_item_id" content="AH0116001" data-react-helmet="true">
				<meta name="description" content="Meta Description">
			</head>
			<body>
				<h1>Pagina de produto 1</h1>
			</body>
		<html>
`))

	httpmock.RegisterResponder("GET", "https://www.teste.com/ah0116002-sandalia-basica-couro-bege-detalhe-laser/p", httpmock.NewStringResponder(200, `
		<!DOCTYPE html>
			<head>
				<title>Titulo do site</title>
				<meta property="product:retailer_item_id" content="AH0116002" data-react-helmet="true">
				<meta name="description" content="Meta Description">
			</head>
			<body>
				<h1>Pagina de produto 2</h1>
			</body>
		<html>
`))

	ecommerce := pkg.NewEcommerce([]string{
		"https://www.teste.com/ah0116001-sandalia-basica-couro-preta-detalhe-laser/p",
		"https://www.teste.com/ah0116002-sandalia-basica-couro-bege-detalhe-laser/p",
	}, 10, true)
	response := ecommerce.Call()

	assert.NotNil(t, response)
	assert.Equal(t, int64(38), response.TotalCaracters)
	assert.Equal(t, 2, len(response.Data))
}
