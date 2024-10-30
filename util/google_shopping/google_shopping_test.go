package googleshopping_test

import (
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"

	googleshopping "github.com/gildemberg-santos/webcrawlerurl_v2/util/google_shopping"
)

func TestGoogleShopping_Call(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "http://www.teste.com", httpmock.NewStringResponder(200, `
	<?xml version="1.0" encoding="UTF-8"?>
	<feed xmlns="http://www.w3.org/2005/Atom">
		<entry>
			<g:link>https://www.google.com/shopping/product/1</g:link>
			<g:image_link>https://www.google.com/shopping/image/product/1.jpg</g:image_link>
		</entry>
		<entry>
			<g:link>https://www.google.com/shopping/product/2</g:link>
			<g:image_link>https://www.google.com/shopping/image/product/2.jpg</g:image_link>
		</entry>
	</feed>
	`))

	googleShopping := googleshopping.NewGoogleShopping("http://www.teste.com", 240)
	err := googleShopping.Call()
	assert.Nil(t, err)
	assert.Equal(t, "http://www.teste.com", googleShopping.UrlLocation)
	assert.Equal(t, "https://www.google.com/shopping/product/1", googleShopping.Feed.Entry[0].Link.Value)
	assert.Equal(t, "https://www.google.com/shopping/product/2", googleShopping.Feed.Entry[1].Link.Value)
	assert.Equal(t, "https://www.google.com/shopping/image/product/1.jpg", googleShopping.Feed.Entry[0].ImageLink.Value)
	assert.Equal(t, "https://www.google.com/shopping/image/product/2.jpg", googleShopping.Feed.Entry[1].ImageLink.Value)
}
