package pkg

import (
	"log"

	googleshopping "github.com/gildemberg-santos/webcrawlerurl_v2/util/google_shopping"
	"github.com/gildemberg-santos/webcrawlerurl_v2/util/normalize"
	"github.com/gildemberg-santos/webcrawlerurl_v2/util/url_match"
)

type EcommerceGoogleShopping struct {
	Visited        map[string]bool          `json:"-"`
	UrlPattern     *url_match.UrlMatch      `json:"-"`
	MaxTimeout     int64                    `json:"-"`
	IsNormalizeUrl bool                     `json:"-"`
	Url            string                   `json:"url,omitempty"`
	Products       []GooogleShoppingProduct `json:"products,omitempty"`
	Urls           []string                 `json:"urls,omitempty"`
}

type GooogleShoppingProduct struct {
	ID           string `json:"id,omitempty"`
	Title        string `json:"title,omitempty"`
	Description  string `json:"description,omitempty"`
	Url          string `json:"url,omitempty"`
	Image        string `json:"image,omitempty"`
	Price        string `json:"price,omitempty"`
	SalePrice    string `json:"sale_price,omitempty"`
	Availability string `json:"availability,omitempty"`
	Condition    string `json:"condition,omitempty"`
	Gender       string `json:"gender,omitempty"`
	Size         string `json:"size,omitempty"`
	AgeGroup     string `json:"age_group,omitempty"`
	Color        string `json:"color,omitempty"`
	Months       string `json:"months,omitempty"`
	Amount       string `json:"amount,omitempty"`
	Downpayment  string `json:"downpayment,omitempty"`
	CreditType   string `json:"credit_type,omitempty"`
}

func NewEcommerceGoogleShopping(url, urlPattern string, maxTimeout int64, isNormalizeUrl bool) *EcommerceGoogleShopping {
	return &EcommerceGoogleShopping{
		Url:            url,
		UrlPattern:     url_match.NewUrlMatch(urlPattern),
		Visited:        map[string]bool{},
		MaxTimeout:     maxTimeout,
		IsNormalizeUrl: isNormalizeUrl,
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
		url_entry := entry.Link.Value

		if s.IsNormalizeUrl {
			url_entry, _ = normalize.NewNormalizeUrl(url_entry).GetUrl()
		}

		if s.Visited[url_entry] {
			continue
		}

		if s.UrlPattern.Call(url_entry) {
			var product GooogleShoppingProduct

			product.ID = entry.ID.Value
			product.Title = entry.Title.Value
			product.Description = entry.Description.Value
			product.Url = url_entry
			product.Image = entry.ImageLink.Value
			product.Price = entry.Price.Value
			product.SalePrice = entry.SalePrice.Value
			product.Availability = entry.Availability.Value
			product.Condition = entry.Condition.Value
			product.Gender = entry.Gender.Value
			product.Size = entry.Size.Value
			product.AgeGroup = entry.AgeGroup.Value
			product.Color = entry.Color.Value
			product.Months = entry.Installment.Months.Value
			product.Amount = entry.Installment.Amount.Value
			product.Downpayment = entry.Installment.Downpayment.Value
			product.CreditType = entry.Installment.CreditType.Value

			s.Urls = append(s.Urls, url_entry)
			s.Products = append(s.Products, product)
		}
		s.Visited[url_entry] = true
	}

	for _, item := range googleShopping.RSS.Item {
		url_item := item.Link.Value

		if s.IsNormalizeUrl {
			url_item, _ = normalize.NewNormalizeUrl(url_item).GetUrl()
		}

		if s.Visited[url_item] {
			continue
		}

		if s.UrlPattern.Call(url_item) {
			var product GooogleShoppingProduct

			product.ID = item.ID.Value
			product.Title = item.Title.Value
			product.Description = item.Description.Value
			product.Url = url_item
			product.Image = item.ImageLink.Value
			product.Price = item.Price.Value
			product.SalePrice = item.SalePrice.Value
			product.Availability = item.Availability.Value
			product.Condition = item.Condition.Value
			product.Gender = item.Gender.Value
			product.Size = item.Size.Value
			product.AgeGroup = item.AgeGroup.Value
			product.Color = item.Color.Value
			product.Months = item.Installment.Months.Value
			product.Amount = item.Installment.Amount.Value
			product.Downpayment = item.Installment.Downpayment.Value
			product.CreditType = item.Installment.CreditType.Value

			s.Urls = append(s.Urls, url_item)
			s.Products = append(s.Products, product)
		}
		s.Visited[url_item] = true
	}

	for _, item := range googleShopping.RDF.Item {
		url_item := item.Link.Value

		if s.IsNormalizeUrl {
			url_item, _ = normalize.NewNormalizeUrl(url_item).GetUrl()
		}

		if s.Visited[url_item] {
			continue
		}

		if s.UrlPattern.Call(url_item) {
			var product GooogleShoppingProduct

			product.ID = item.ID.Value
			product.Title = item.Title.Value
			product.Description = item.Description.Value
			product.Url = url_item
			product.Image = item.ImageLink.Value
			product.Price = item.Price.Value
			product.SalePrice = item.SalePrice.Value
			product.Availability = item.Availability.Value
			product.Condition = item.Condition.Value
			product.Gender = item.Gender.Value
			product.Size = item.Size.Value
			product.AgeGroup = item.AgeGroup.Value
			product.Color = item.Color.Value
			product.Months = item.Installment.Months.Value
			product.Amount = item.Installment.Amount.Value
			product.Downpayment = item.Installment.Downpayment.Value
			product.CreditType = item.Installment.CreditType.Value

			s.Urls = append(s.Urls, url_item)
			s.Products = append(s.Products, product)
		}
		s.Visited[url_item] = true
	}

	return nil
}
