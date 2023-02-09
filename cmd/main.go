package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gildemberg-santos/webcrawlerurl_v2/pkg"
)

func main() {
	http.HandleFunc("/chatgpt3", ChatGpt3)
	http.ListenAndServe(":8080", nil)
}

type Response struct {
	Title       string `json:"title"`
	Paragraph   string `json:"paragraph"`
	Description string `json:"description"`
}

func ChatGpt3(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	url := r.URL.Query().Get("url")

	pagina := pkg.LoadPage{
		Url: url,
	}

	err := pagina.Load()
	if err != nil {
		fmt.Println("Erro ao acessar a URL:", err)
		return
	}

	informatin := pkg.ExtractInformation{}
	informatin.Init(pagina.Source, pagina.Url, 3, 3, 3)
	informatin.Call()

	informatinResponse := Response{
		Title:       informatin.MainTitle,
		Paragraph:   informatin.MainParagraph,
		Description: informatin.MetaDescription,
	}

	json.NewEncoder(w).Encode(informatinResponse)
}
