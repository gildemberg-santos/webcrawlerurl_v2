package sitemap_test

import (
	"log"
	"testing"

	sitemap "github.com/gildemberg-santos/webcrawlerurl_v2/util/site_map"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestSiteMap_Call(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "http://www.teste.com", httpmock.NewStringResponder(200, `
	<?xml version="1.0" encoding="UTF-8"?>
	<urlset
				xmlns="http://www.sitemaps.org/schemas/sitemap/0.9"
				xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
				xsi:schemaLocation="http://www.sitemaps.org/schemas/sitemap/0.9
							http://www.sitemaps.org/schemas/sitemap/0.9/sitemap.xsd">
		<url>
			<loc>https://leadster.com.br/</loc>
			<lastmod>2022-11-08T18:29:20+00:00</lastmod>
			<priority>1.00</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/leads-qualificados-com-anuncios/</loc>
			<lastmod>2022-11-08T18:29:21+00:00</lastmod>
			<priority>0.80</priority>
		</url>
	</urlset>
	`))

	siteMap := sitemap.NewSiteMap("http://www.teste.com")
	err := siteMap.Call()
	log.Println(err)
	assert.Nil(t, err)
}
