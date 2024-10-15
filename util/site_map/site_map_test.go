package sitemap_test

import (
	"testing"

	sitemap "github.com/gildemberg-santos/webcrawlerurl_v2/util/site_map"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestSiteMap_Call(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://leadster.com.br", httpmock.NewStringResponder(200, `
	<?xml version="1.0" encoding="UTF-8"?>
	<urlset
		xmlns="http://www.sitemaps.org/schemas/sitemap/0.9"
		xmlns:xhtml="http://www.w3.org/1999/xhtml"
		xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
		xsi:schemaLocation="http://www.sitemaps.org/schemas/sitemap/0.9
		http://www.sitemaps.org/schemas/sitemap/0.9/sitemap.xsd">
		<url>
			<loc>https://leadster.com.br/pt</loc>
			<lastmod>2022-11-08T18:29:20+00:00</lastmod>
			<priority>1.00</priority>
			<xhtml:link rel="alternate" hreflang="pt" href="https://leadster.com.br/pt"/>
			<xhtml:link rel="alternate" hreflang="en" href="https://leadster.com.br/en"/>
		</url>
		<url>
			<loc>https://leadster.com.br/pt/leads-qualificados-com-anuncios/</loc>
			<lastmod>2022-11-08T18:29:21+00:00</lastmod>
			<priority>0.80</priority>
			<xhtml:link rel="alternate" hreflang="pt" href="https://leadster.com.br/pt/leads-qualificados-com-anuncios/"/>
			<xhtml:link rel="alternate" hreflang="en" href="https://leadster.com.br/en/leads-qualificados-com-anuncios/"/>
		</url>
	</urlset>
	`))

	siteMap := sitemap.NewSiteMap("https://leadster.com.br")
	err := siteMap.Call()
	assert.Nil(t, err)
	assert.Equal(t, "https://leadster.com.br", siteMap.UrlLocation)
	assert.Equal(t, "https://leadster.com.br/pt", siteMap.Urlset.URLs[0].Loc)
	assert.Equal(t, "https://leadster.com.br/pt", siteMap.Urlset.URLs[0].Link[0].Href)
	assert.Equal(t, "https://leadster.com.br/en", siteMap.Urlset.URLs[0].Link[1].Href)
	assert.Equal(t, "https://leadster.com.br/pt/leads-qualificados-com-anuncios/", siteMap.Urlset.URLs[1].Loc)
	assert.Equal(t, "https://leadster.com.br/pt/leads-qualificados-com-anuncios/", siteMap.Urlset.URLs[1].Link[0].Href)
	assert.Equal(t, "https://leadster.com.br/en/leads-qualificados-com-anuncios/", siteMap.Urlset.URLs[1].Link[1].Href)
	assert.Len(t, siteMap.Urlset.URLs, 2)
}
