package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kubosho/sandbox-go/treasure/http-server/scraper"
)

type Page struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	if url == "" {
		http.Error(w, "URL not specified", http.StatusBadRequest)
		return
	}
	content, err := scraper.Get(url)
	if err != nil {
		panic(err)
	}

	page := &Page{Title: content.Title, Description: content.Description}

	w.Header().Set("Content-Type", "application/json")

	enc := json.NewEncoder(w)
	if err := enc.Encode(page); err != nil {
		http.Error(w, "encoding failed", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
