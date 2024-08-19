package main

import (
	"log"

	"github.com/gildemberg-santos/webcrawlerurl_v2/pkg"
	"github.com/gildemberg-santos/webcrawlerurl_v2/util/file"
)

func main() {
	sitemap := struct {
		Url               string
		UrlPattern        string
		UrlSiteMapPattern string
	}{
		Url:               "https://www.usaflex.com.br/sitemap.xml",
		UrlPattern:        "https://www.usaflex.com.br/ac**/p",
		UrlSiteMapPattern: "",
	}

	log.Println("Starting")
	log.Println("Loading sitemap -> ", sitemap.Url)
	urls := pkg.NewEcommerceSitemap(sitemap.Url, sitemap.UrlPattern, sitemap.UrlSiteMapPattern).Call()
	fileJson := file.NewFileJson("urls.json", urls)
	fileJson.Save()
	log.Println("Loading data")
	data := pkg.NewEcommerce(urls.Urls, 1_000_000_000, true).Call()
	fileJson = file.NewFileJson("data.json", data)
	fileJson.Save()
	log.Println("Completed")
}
