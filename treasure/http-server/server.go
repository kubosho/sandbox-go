package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/kubosho/sandbox-go/treasure/http-server/scraper"
)

type Page struct {
	URL         string `json:"url"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

const separator = ","

func fetchPageData(url string) (*Page, error) {
	content, err := scraper.Get(url)
	if err != nil {
		return nil, err
	}

	return &Page{URL: url, Title: content.Title, Description: content.Description}, nil
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	if url == "" {
		http.Error(w, "URL not specified", http.StatusBadRequest)
		return
	}

	if strings.Contains(url, separator) {
		urls := strings.Split(url, separator)
		var pageData []*Page

		for _, u := range urls {
			page, err := fetchPageData(u)
			if err != nil {
				http.Error(w, "page fetch are failed", http.StatusInternalServerError)
				return
			}

			pageData = append(pageData, page)
		}

		w.Header().Set("Content-Type", "application/json")

		enc := json.NewEncoder(w)
		if err := enc.Encode(pageData); err != nil {
			http.Error(w, "encoding failed", http.StatusInternalServerError)
			return
		}
		return
	}

	page, err := fetchPageData(url)
	if err != nil {
		http.Error(w, "page fetch are failed", http.StatusInternalServerError)
		return
	}

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
