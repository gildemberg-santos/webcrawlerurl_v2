package main

import (
	"log"

	"github.com/gildemberg-santos/webcrawlerurl_v2/pkg"
	"github.com/gildemberg-santos/webcrawlerurl_v2/util/file"
	"github.com/gildemberg-santos/webcrawlerurl_v2/util/normalize"
)

// main is the entry point of the program.
//
// It initializes the variables url_base, maxUrlLimit, and maxChunckLimit with default values.
// Then, it creates a new instance of LeadsterAI with the given parameters.
// After that, it calls the Call method of the LeadsterAI instance to start crawling the URLs.
// Finally, it logs the length of the Data field of the LeadsterAI instance and saves the data to a JSON file named "data.json".
func main() {

	body := struct {
		Url         string `json:"url"`
		MaxUrlLimit int64  `json:"max_url_limit"`
		MaxTimeout  int64  `json:"max_timeout"`
		UrlPattern  string `json:"url_pattern"`
		IsLoadFast  bool   `json:"is_load_fast"`
		IsSiteMap   bool   `json:"is_sitemap"`
		IsComplete  bool   `json:"is_complete"`
	}{}

	body.Url = "https://leadster.com.br"
	body.UrlPattern = "https://leadster.com.br**"
	body.MaxUrlLimit = 1
	body.MaxTimeout = 30
	body.IsLoadFast = false
	body.IsSiteMap = false
	body.IsComplete = false

	log.Println("Starting crawler...")

	body.Url, _ = normalize.NewNormalizeUrl(body.Url).GetUrl()
	body.UrlPattern, _ = normalize.NewNormalizeUrl(body.UrlPattern).GetUrl()

	leadsterAI := pkg.NewLeadsterAI(body.Url, body.UrlPattern, body.MaxUrlLimit, body.MaxTimeout, body.IsLoadFast)
	leadsterAI.Call(body.IsSiteMap, body.IsComplete)
	log.Printf("Saving data to file data.json total urls: %d\n", len(leadsterAI.Data))

	fileJson := file.NewFileJson("data.json", leadsterAI)
	fileJson.Save()
	log.Println("Data saved to file data.json")
}
