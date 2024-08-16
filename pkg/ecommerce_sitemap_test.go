package pkg_test

import (
	"testing"

	"github.com/gildemberg-santos/webcrawlerurl_v2/pkg"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestEcommerceSitemap_Call(t *testing.T) {
	httpmock.Activate()

	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://www.teste.com/sitemap.xml", httpmock.NewStringResponder(200, `
		<?xml version="1.0" encoding="UTF-8"?>
		<sitemapindex>
			<sitemap>
				<loc>https://www.teste.com/sitemap/product-0.xml</loc>
				<lastmod>2024-08-15T18:04:23.316Z</lastmod>
			</sitemap>
			<sitemap>
				<loc>https://www.teste.com/sitemap/product-1.xml</loc>
				<lastmod>2024-08-15T18:04:23.316Z</lastmod>
			</sitemap>
		</sitemapindex>
	`))

	httpmock.RegisterResponder("GET", "https://www.teste.com/sitemap/product-0.xml", httpmock.NewStringResponder(200, `
		<?xml version="1.0" encoding="UTF-8"?>
		<urlset
			xmlns="http://www.sitemaps.org/schemas/sitemap/0.9"
			xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
			xsi:schemaLocation="http://www.sitemaps.org/schemas/sitemap/0.9 http://www.sitemaps.org/schemas/sitemap/0.9/sitemap.xsd">
			<url>
				<loc>https://www.teste.com/ag1622005-scarpin-marrom-telha-salto-bloco-couro-enfeite/p</loc>
				<lastmod>2024-08-15T18:08:10.653Z</lastmod>
			</url>
			<url>
				<loc>https://www.teste.com/ag1623001-scarpin-preto-salto-bloco-couro-monograma/p</loc>
				<lastmod>2024-08-15T18:08:10.653Z</lastmod>
			</url>
		</urlset>
	`))

	httpmock.RegisterResponder("GET", "https://www.teste.com/sitemap/product-1.xml", httpmock.NewStringResponder(200, `
		<?xml version="1.0" encoding="UTF-8"?>
		<urlset
			xmlns="http://www.sitemaps.org/schemas/sitemap/0.9"
			xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
			xsi:schemaLocation="http://www.sitemaps.org/schemas/sitemap/0.9 http://www.sitemaps.org/schemas/sitemap/0.9/sitemap.xsd">
			<url>
				<loc>https://www.teste.com/ah0116001-sandalia-basica-couro-preta-detalhe-laser/p</loc>
				<lastmod>2024-08-15T18:08:11.912Z</lastmod>
			</url>
			<url>
				<loc>https://www.teste.com/ah0116002-sandalia-basica-couro-bege-detalhe-laser/p</loc>
				<lastmod>2024-08-15T18:08:11.912Z</lastmod>
			</url>
		</urlset>
`))

	ecommerceSitemap := pkg.NewEcommerceSitemap("https://www.teste.com/sitemap.xml", "", "")
	response := ecommerceSitemap.Call()

	assert.NotNil(t, response)
	assert.Equal(t, "https://www.teste.com/sitemap.xml", response.Url)
	assert.Equal(t, 2, len(response.UrlSiteMap))
	assert.Equal(t, "https://www.teste.com/sitemap/product-0.xml", response.UrlSiteMap[0])
	assert.Equal(t, "https://www.teste.com/sitemap/product-1.xml", response.UrlSiteMap[1])
	assert.Equal(t, 4, len(response.Urls))
	assert.Equal(t, "https://www.teste.com/ag1622005-scarpin-marrom-telha-salto-bloco-couro-enfeite/p", response.Urls[0])
	assert.Equal(t, "https://www.teste.com/ag1623001-scarpin-preto-salto-bloco-couro-monograma/p", response.Urls[1])
	assert.Equal(t, "https://www.teste.com/ah0116001-sandalia-basica-couro-preta-detalhe-laser/p", response.Urls[2])
	assert.Equal(t, "https://www.teste.com/ah0116002-sandalia-basica-couro-bege-detalhe-laser/p", response.Urls[3])
}
