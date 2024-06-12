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
	http.HandleFunc("/leadster_ai", RouteLeadsterAI)

	log.Println("Listening on port " + GetPort())
	err := http.ListenAndServe(GetPort(), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func RouteSmartCall(w http.ResponseWriter, r *http.Request) {
	log.Println("RouteSmartCall")
	w.Header().Set("Content-Type", "application/json")
	url, _ := normalize.NewNormalizeUrl(r.URL.Query().Get("url")).GetUrl()

	smartCall := pkg.NewSmartCall(url)

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

	mapping := pkg.NewMappingUrl(url, limit, nil)

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
	limitchunck, _ := strconv.ParseInt(r.URL.Query().Get("limitchunck"), 10, 64)
	log.Default().Println("limitchunck -> ", limitchunck)

	readtext := pkg.NewReadText(url, int64(limitchunck), nil)

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
	w.Header().Set("Content-Type", "application/json")

	url, _ := normalize.NewNormalizeUrl(r.URL.Query().Get("url")).GetUrl()
	maxUrlLimit, _ := strconv.ParseInt(r.URL.Query().Get("maxurllimit"), 10, 64)
	maxChunckLimit, _ := strconv.ParseInt(r.URL.Query().Get("maxchuncklimit"), 10, 64)

	leadsterAI := pkg.NewLeadsterAI(url, maxUrlLimit, maxChunckLimit)

	response := leadsterAI.Call()

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
