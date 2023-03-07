package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gildemberg-santos/webcrawlerurl_v2/pkg"
)

func main() {
	log.Println("Starting server...")
	http.HandleFunc("/chatgpt3", RouteChatGpt3)
	http.HandleFunc("/mappingurl", RouteMappingUrl)

	log.Println("Listening on port " + GetPort())
	err := http.ListenAndServe(GetPort(), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func RouteChatGpt3(w http.ResponseWriter, r *http.Request) {
	log.Println("RouteChatGpt3")
	w.Header().Set("Content-Type", "application/json")
	uri := pkg.NormalizeUrl{Url: r.URL.Query().Get("url")}
	url, _ := uri.GetUrl()

	chatgpt3 := pkg.ChatGpt3{
		Url: url,
	}

	response, err := chatgpt3.Call()
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
	uri := pkg.NormalizeUrl{Url: r.URL.Query().Get("url")}
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	url, _ := uri.GetUrl()

	mapping := pkg.MappingUrl{
		Url:   url,
		Limit: int(limit),
	}

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

func GetPort() string {
	var port = os.Getenv("PORT")

	if port == "" {
		port = "4747"
		log.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}

	return ":" + port
}
