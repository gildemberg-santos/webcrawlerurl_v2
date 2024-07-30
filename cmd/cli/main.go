package main

import (
	"log"

	"github.com/gildemberg-santos/webcrawlerurl_v2/pkg"
	"github.com/gildemberg-santos/webcrawlerurl_v2/util/file"
)

// main is the entry point of the program.
//
// It initializes the variables url_base, maxUrlLimit, and maxChunckLimit with default values.
// Then, it creates a new instance of LeadsterAI with the given parameters.
// After that, it calls the Call method of the LeadsterAI instance to start crawling the URLs.
// Finally, it logs the length of the Data field of the LeadsterAI instance and saves the data to a JSON file named "data.json".
func main() {
	var url_base string = "https://www.usaflex.com.br/sitemap/product-1.xml"
	var url_pattern string = "https://www.usaflex.com.br/**"
	var maxUrlLimit int64 = 100
	var maxChunckLimit int64 = 2_000_000
	var maxCaracterLimit int64 = 2_000_000
	var loadPageFast bool = false

	log.Println("Starting crawler...")
	leadsterAI := pkg.NewLeadsterAI(url_base, maxUrlLimit, maxChunckLimit, maxCaracterLimit, url_pattern, loadPageFast)
	leadsterAI.Call(true, false)
	log.Printf("Saving data to file data.json total urls: %d\n", len(leadsterAI.Data))

	fileJson := file.NewFileJson("data.json", leadsterAI)
	fileJson.Save()
	log.Println("Data saved to file data.json")
}
