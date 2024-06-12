package main

import (
	"log"

	"github.com/gildemberg-santos/webcrawlerurl_v2/pkg"
	"github.com/gildemberg-santos/webcrawlerurl_v2/util/file"
)

func main() {
	var url_base string = "https://leadster.com.br/"
	var maxUrlLimit int64 = 2_000_000
	var maxChunckLimit int64 = 2_000_000

	leadsterAI := pkg.NewLeadsterAI(url_base, maxUrlLimit, maxChunckLimit)
	leadsterAI.Call()

	log.Println(len(leadsterAI.Data))

	fileJson := file.NewFileJson("data.json", leadsterAI)
	fileJson.Save()
}
