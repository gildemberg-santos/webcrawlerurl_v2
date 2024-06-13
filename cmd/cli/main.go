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
	var url_base string = "https://leadster.com.br/"
	var maxUrlLimit int64 = 28
	var maxChunckLimit int64 = 2_000_000
	var maxCaracterLimit int64 = 2_000_000

	leadsterAI := pkg.NewLeadsterAI(url_base, maxUrlLimit, maxChunckLimit, maxCaracterLimit)
	leadsterAI.Call()

	log.Println(len(leadsterAI.Data))

	fileJson := file.NewFileJson("data.json", leadsterAI)
	fileJson.Save()
}
