package main

import (
	"encoding/json"
	"net/http"

	"github.com/gildemberg-santos/webcrawlerurl_v2/pkg"
)

func main() {
	http.HandleFunc("/chatgpt3", ChatGpt3)
	http.HandleFunc("/urls", Urls)
	http.ListenAndServe(":8080", nil)
}

type ResponseChatGpt3 struct {
	ChatGpt struct {
		Title       string `json:"main_header"`
		Paragraph   string `json:"main_paragraph"`
		Description string `json:"meta_description"`
	} `json:"chatgpt"`
	Url        string  `json:"url"`
	Timestamp  float64 `json:"ts"`
	Scone      float32 `json:"scone"`
	StatusCode int     `json:"status_code"`
}

type ResponseUrls struct {
	Urls       []string `json:"urls"`
	Url        string   `json:"url"`
	Timestamp  float64  `json:"ts"`
	StatusCode int      `json:"status_code"`
}

type ResponseErro struct {
	Erro       string  `json:"erro"`
	Timestamp  float64 `json:"ts"`
	StatusCode int     `json:"status_code"`
}

func ChatGpt3(w http.ResponseWriter, r *http.Request) {
	time := pkg.Timestamp{}
	time.Start()
	w.Header().Set("Content-Type", "application/json")

	url := r.URL.Query().Get("url")
	if url == "" {
		time.End()
		erroResponse := ResponseErro{
			Erro:       "url is empty",
			Timestamp:  time.GetTime(),
			StatusCode: 500,
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(erroResponse)
		return
	}

	pagina := pkg.LoadPage{
		Url: r.URL.Query().Get("url"),
	}

	err := pagina.Load()
	if err != nil {
		time.End()
		erroResponse := ResponseErro{
			Erro:       err.Error(),
			Timestamp:  time.GetTime(),
			StatusCode: pagina.StatusCode,
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(erroResponse)
		return
	}

	informatin := pkg.ExtractInformation{}
	informatin.Init(pagina.Source, pagina.Url, 3, 5, 30)
	informatin.Call()
	time.End()

	score := pkg.Score{}
	score.Init(pagina.Url, &informatin)
	score.Call()

	informatinResponse := ResponseChatGpt3{
		Url:        informatin.Url,
		Timestamp:  time.GetTime(),
		Scone:      score.GetScore(),
		StatusCode: pagina.StatusCode,
	}

	informatinResponse.ChatGpt.Title = informatin.MainTitle
	informatinResponse.ChatGpt.Paragraph = informatin.MainParagraph
	informatinResponse.ChatGpt.Description = informatin.MetaDescription

	json.NewEncoder(w).Encode(informatinResponse)
}

func Urls(w http.ResponseWriter, r *http.Request) {
	time := pkg.Timestamp{}
	time.Start()
	w.Header().Set("Content-Type", "application/json")
	url := r.URL.Query().Get("url")
	time.End()

	urlsResponse := ResponseUrls{
		Urls:       []string{""},
		Url:        url,
		Timestamp:  time.GetTime(),
		StatusCode: 200,
	}

	json.NewEncoder(w).Encode(urlsResponse)
}
