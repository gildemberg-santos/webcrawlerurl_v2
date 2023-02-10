package main

import (
	"encoding/json"
	"net/http"

	"github.com/gildemberg-santos/webcrawlerurl_v2/pkg"
)

func main() {
	http.HandleFunc("/chatgpt3", RouteChatGpt3)
	http.ListenAndServe(":8080", nil)
}

func RouteChatGpt3(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	chatgpt3 := pkg.ChatGpt3{
		Url: r.URL.Query().Get("url"),
	}

	response, err := chatgpt3.Call("teste")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
