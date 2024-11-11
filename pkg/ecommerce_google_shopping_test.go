package pkg_test

import (
	"testing"

	"github.com/gildemberg-santos/webcrawlerurl_v2/pkg"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestEcommerceGoogleShopping_Call(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "http://www.teste.com/google_shopping.xml", httpmock.NewStringResponder(200, `
	<?xml version="1.0" encoding="UTF-8"?>
	<feed xmlns="http://www.w3.org/2005/Atom">
		<entry>
			<g:id>1</g:id>
			<g:title>Product 1</g:title>
			<g:description>Description product 1</g:description>
			<g:link>https://www.google.com/shopping/product/1</g:link>
			<g:image_link>https://www.google.com/shopping/image/product/1.jpg</g:image_link>
			<g:price>5.00</g:price>
			<g:sale_price>4.00</g:sale_price>
			<g:availability>in stock</g:availability>
			<g:condition>new</g:condition>
			<g:gender>mole</g:gender>
			<g:size>size</g:size>
			<g:age_group>age group</g:age_group>
			<g:color>color</g:color>
			<g:installment>
				<g:months>12</g:months>
				<g:amount>5.00</g:amount>
				<g:downpayment>USD</g:downpayment>
				<g:credit_type>monthly</g:credit_type>
			</g:installment>
		</entry>
		<entry>
			<g:id>2</g:id>
			<g:title>Product 2</g:title>
			<g:description>Description product 2</g:description>
			<g:link>https://www.google.com/shopping/product/2</g:link>
			<g:image_link>https://www.google.com/shopping/image/product/2.jpg</g:image_link>
			<g:price>10.00</g:price>
			<g:sale_price>8.00</g:sale_price>
			<g:availability>in stock</g:availability>
			<g:condition>new</g:condition>
			<g:gender>male</g:gender>
			<g:size>size</g:size>
			<g:age_group>age group</g:age_group>
			<g:color>color</g:color>
			<g:installment>
				<g:months>12</g:months>
				<g:amount>5.00</g:amount>
				<g:downpayment>USD</g:downpayment>
				<g:credit_type>monthly</g:credit_type>
			</g:installment>
		</entry>
	</feed>
	`))

	ecommerceGoogleShopping := pkg.NewEcommerceGoogleShopping("http://www.teste.com/google_shopping.xml", "https://www.google.com/shopping/product/**", 30)
	response := ecommerceGoogleShopping.Call()

	assert.NotNil(t, response)
	assert.Equal(t, "http://www.teste.com/google_shopping.xml", ecommerceGoogleShopping.Url)

	assert.Len(t, ecommerceGoogleShopping.Urls, 2)
	assert.Equal(t, "https://www.google.com/shopping/product/1", ecommerceGoogleShopping.Urls[0])
	assert.Equal(t, "https://www.google.com/shopping/product/2", ecommerceGoogleShopping.Urls[1])

	assert.Len(t, ecommerceGoogleShopping.Products, 2)
	assert.Equal(t, "1", ecommerceGoogleShopping.Products[0].ID)
	assert.Equal(t, "Product 1", ecommerceGoogleShopping.Products[0].Title)
	assert.Equal(t, "Description product 1", ecommerceGoogleShopping.Products[0].Description)
	assert.Equal(t, "https://www.google.com/shopping/product/1", ecommerceGoogleShopping.Products[0].Url)
	assert.Equal(t, "https://www.google.com/shopping/image/product/1.jpg", ecommerceGoogleShopping.Products[0].Image)
	assert.Equal(t, "5.00", ecommerceGoogleShopping.Products[0].Price)
	assert.Equal(t, "4.00", ecommerceGoogleShopping.Products[0].SalePrice)
	assert.Equal(t, "in stock", ecommerceGoogleShopping.Products[0].Availability)
	assert.Equal(t, "new", ecommerceGoogleShopping.Products[0].Condition)
	assert.Equal(t, "mole", ecommerceGoogleShopping.Products[0].Gender)
	assert.Equal(t, "size", ecommerceGoogleShopping.Products[0].Size)
	assert.Equal(t, "age group", ecommerceGoogleShopping.Products[0].AgeGroup)
	assert.Equal(t, "color", ecommerceGoogleShopping.Products[0].Color)
	assert.Equal(t, "12", ecommerceGoogleShopping.Products[0].Months)
	assert.Equal(t, "5.00", ecommerceGoogleShopping.Products[0].Amount)
	assert.Equal(t, "USD", ecommerceGoogleShopping.Products[0].Downpayment)
	assert.Equal(t, "monthly", ecommerceGoogleShopping.Products[0].CreditType)

	assert.Equal(t, "2", ecommerceGoogleShopping.Products[1].ID)
	assert.Equal(t, "Product 2", ecommerceGoogleShopping.Products[1].Title)
	assert.Equal(t, "Description product 2", ecommerceGoogleShopping.Products[1].Description)
	assert.Equal(t, "https://www.google.com/shopping/product/2", ecommerceGoogleShopping.Products[1].Url)
	assert.Equal(t, "https://www.google.com/shopping/image/product/2.jpg", ecommerceGoogleShopping.Products[1].Image)
	assert.Equal(t, "10.00", ecommerceGoogleShopping.Products[1].Price)
	assert.Equal(t, "8.00", ecommerceGoogleShopping.Products[1].SalePrice)
	assert.Equal(t, "in stock", ecommerceGoogleShopping.Products[1].Availability)
	assert.Equal(t, "new", ecommerceGoogleShopping.Products[1].Condition)
	assert.Equal(t, "male", ecommerceGoogleShopping.Products[1].Gender)
	assert.Equal(t, "size", ecommerceGoogleShopping.Products[1].Size)
	assert.Equal(t, "age group", ecommerceGoogleShopping.Products[1].AgeGroup)
	assert.Equal(t, "color", ecommerceGoogleShopping.Products[1].Color)
	assert.Equal(t, "12", ecommerceGoogleShopping.Products[1].Months)
	assert.Equal(t, "5.00", ecommerceGoogleShopping.Products[1].Amount)
	assert.Equal(t, "USD", ecommerceGoogleShopping.Products[1].Downpayment)
	assert.Equal(t, "monthly", ecommerceGoogleShopping.Products[1].CreditType)
}
