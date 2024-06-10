package serve

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
	http.HandleFunc("/chatgpt3", RouteChatGpt3)
	http.HandleFunc("/mappingurl", RouteMappingUrl)
	http.HandleFunc("/readtext", RouteReadText)

	log.Println("Listening on port " + GetPort())
	err := http.ListenAndServe(GetPort(), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func RouteChatGpt3(w http.ResponseWriter, r *http.Request) {
	log.Println("RouteChatGpt3")
	w.Header().Set("Content-Type", "application/json")
	url, _ := normalize.NewNormalizeUrl(r.URL.Query().Get("url")).GetUrl()

	chatgpt3 := pkg.NewChatGpt3(url)

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
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	url, _ := normalize.NewNormalizeUrl(r.URL.Query().Get("url")).GetUrl()

	mapping := pkg.NewMappingUrl(url, int(limit))

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

	readtext := pkg.NewReadText(url)

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

func GetPort() string {
	var port = os.Getenv("PORT")

	if port == "" {
		port = "4747"
		log.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}

	return ":" + port
}
