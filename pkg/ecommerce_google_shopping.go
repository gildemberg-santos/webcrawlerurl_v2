package pkg

import (
	"log"

	googleshopping "github.com/gildemberg-santos/webcrawlerurl_v2/util/google_shopping"
	"github.com/gildemberg-santos/webcrawlerurl_v2/util/url_match"
)

type EcommerceGoogleShopping struct {
	Visited    map[string]bool          `json:"-"`
	UrlPattern *url_match.UrlMatch      `json:"-"`
	MaxTimeout int64                    `json:"-"`
	Url        string                   `json:"url,omitempty"`
	Products   []GooogleShoppingProduct `json:"products,omitempty"`
	Urls       []string                 `json:"urls,omitempty"`
}

type GooogleShoppingProduct struct {
	ID           string `json:"id,omitempty"`
	Title        string `json:"title,omitempty"`
	Description  string `json:"description,omitempty"`
	Url          string `json:"url,omitempty"`
	Image        string `json:"image,omitempty"`
	Price        string `json:"price,omitempty"`
	Availability string `json:"availability,omitempty"`
}

func NewEcommerceGoogleShopping(url, urlPattern string, maxTimeout int64) *EcommerceGoogleShopping {
	return &EcommerceGoogleShopping{
		Url:        url,
		UrlPattern: url_match.NewUrlMatch(urlPattern),
		Visited:    map[string]bool{},
		MaxTimeout: maxTimeout,
	}
}

func (s *EcommerceGoogleShopping) Call() *EcommerceGoogleShopping {
	if err := s.crawler(s.Url); err != nil {
		log.Default().Println(err)
		return s
	}

	return s
}

func (s *EcommerceGoogleShopping) crawler(url string) error {
	s.Visited[url] = true

	googleShopping := googleshopping.NewGoogleShopping(url, s.MaxTimeout)

	if err := googleShopping.Call(); err != nil {
		log.Default().Println("Error request google shopping: ", err)
		return err
	}

	for _, entry := range googleShopping.Feed.Entry {
		if s.Visited[entry.Link.Value] {
			continue
		}
		if s.UrlPattern.Call(entry.Link.Value) {
			var product GooogleShoppingProduct

			product.ID = entry.ID.Value
			product.Title = entry.Title.Value
			product.Description = entry.Description.Value
			product.Url = entry.Link.Value
			product.Image = entry.ImageLink.Value
			product.Price = entry.Price.Value
			product.Availability = entry.Availability.Value

			s.Urls = append(s.Urls, entry.Link.Value)
			s.Products = append(s.Products, product)

			log.Println("Product found: ", product)
		}
		s.Visited[entry.Link.Value] = true
	}

	return nil
}
