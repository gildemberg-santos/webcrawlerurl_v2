package cli

import (
	"log"

	"github.com/gildemberg-santos/webcrawlerurl_v2/pkg"
	"github.com/gildemberg-santos/webcrawlerurl_v2/util/normalize"
)

func main() {
	limit := 28
	url, _ := normalize.NewNormalizeUrl("https://leadster.com.br/").GetUrl()

	mapping := pkg.MappingUrl{
		Url:   url,
		Limit: limit,
	}

	response, err := mapping.Call()
	if err != nil {
		log.Println("Error")
		log.Println(response)
		return
	}

	log.Println("Success")
	log.Println(response)
}
