package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gildemberg-santos/webcrawlerurl_v2/pkg"
	"github.com/gildemberg-santos/webcrawlerurl_v2/util/normalize"
)

func main() {
	log.Println("Starting server...")
	http.HandleFunc("/smart_call", RouteSmartCall)
	http.HandleFunc("/mappingurl", RouteMappingUrl)
	http.HandleFunc("/readtext", RouteReadText)
	http.HandleFunc("/crawler_leadster_ai", RouteLeadsterAI)

	log.Println("Listening on port " + GetPort())
	err := http.ListenAndServe(GetPort(), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func RouteSmartCall(w http.ResponseWriter, r *http.Request) {
	log.Println("RouteSmartCall")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	url, _ := normalize.NewNormalizeUrl(r.URL.Query().Get("url")).GetUrl()

	smartCall := pkg.NewSmartCall(url, true)

	response, err := smartCall.Call()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	log.Println("Success")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func RouteMappingUrl(w http.ResponseWriter, r *http.Request) {
	log.Println("RouteMappingUrl")
	w.Header().Set("Content-Type", "application/json")
	limit, _ := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 64)

	url, _ := normalize.NewNormalizeUrl(r.URL.Query().Get("url")).GetUrl()

	mapping := pkg.NewMappingUrl(url, limit, true, nil)

	response, err := mapping.Call()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	log.Println("Success")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func RouteReadText(w http.ResponseWriter, r *http.Request) {
	log.Println("RouteReadText")
	w.Header().Set("Content-Type", "application/json")

	url, _ := normalize.NewNormalizeUrl(r.URL.Query().Get("url")).GetUrl()

	readtext := pkg.NewReadText(url, nil, true)

	response, err := readtext.Call()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	log.Println("Success")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func RouteLeadsterAI(w http.ResponseWriter, r *http.Request) {
	log.Println("RouteLeadsterAI")
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

	log.Println("Body: ", body)

	body.Url, _ = normalize.NewNormalizeUrl(body.Url).GetUrl()

	if body.UrlPattern != "" {
		body.UrlPattern, _ = normalize.NewNormalizeUrl(body.UrlPattern).GetUrl()
	}

	leadsterAI := pkg.NewLeadsterAI(body.Url, body.UrlPattern, body.MaxUrlLimit, body.MaxTimeout, body.IsLoadFast)
	response := leadsterAI.Call(body.IsSiteMap, body.IsComplete)

	log.Println("Success")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func GetPort() string {
	var port = os.Getenv("PORT")

	if port == "" {
		port = "4747"
		log.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}

	return ":" + port
}
