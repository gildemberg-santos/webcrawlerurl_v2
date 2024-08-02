package helloworld

import (
	"encoding/json"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/gildemberg-santos/webcrawlerurl_v2/pkg"
	"github.com/gildemberg-santos/webcrawlerurl_v2/util/normalize"
)

func init() {
	functions.HTTP("Call", RouteLeadsterAI)
}

func RouteLeadsterAI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	body := struct {
		Url         string `json:"url"`
		UrlPattern  string `json:"url_pattern"`
		MaxUrlLimit int64  `json:"max_url_limit"`
		MaxTimeout  int64  `json:"max_timeout"`
		IsLoadFast  bool   `json:"is_load_fast"`
		IsSiteMap   bool   `json:"is_sitemap"`
		IsComplete  bool   `json:"is_complete"`
	}{}

	json.NewDecoder(r.Body).Decode(&body)

	body.Url, _ = normalize.NewNormalizeUrl(body.Url).GetUrl()
	body.UrlPattern, _ = normalize.NewNormalizeUrl(body.UrlPattern).GetUrl()

	leadsterAI := pkg.NewLeadsterAI(body.Url, body.UrlPattern, body.MaxUrlLimit, body.MaxTimeout, body.IsLoadFast)
	response := leadsterAI.Call(body.IsSiteMap, body.IsComplete)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
