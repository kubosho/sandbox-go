package scraper

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

type Page struct {
	Title       string
	Description string
	OgImage     string
	OgTitle     string
}

func isDescription(attrs []html.Attribute) bool {
	for _, attr := range attrs {
		if attr.Key == "name" && attr.Val == "description" {
			return true
		}
	}
	return false
}

func isOgTitle(attrs []html.Attribute) bool {
	for _, attr := range attrs {
		if attr.Key == "property" && attr.Val == "og:title" {
			return true
		}
	}
	return false
}

func isOgImage(attrs []html.Attribute) bool {
	for _, attr := range attrs {
		if attr.Key == "property" && attr.Val == "og:image" {
			return true
		}
	}
	return false
}

func Get(url string) (*Page, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var f func(*html.Node)
	var title string
	var desc string
	var ogImage string
	var ogTitle string

	f = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "title" {
			title = node.FirstChild.Data
		}

		if node.Type == html.ElementNode && node.Data == "meta" {
			if isDescription(node.Attr) {
				for _, attr := range node.Attr {
					desc = attr.Val
				}
			}

			if isOgTitle(node.Attr) {
				for _, attr := range node.Attr {
					ogTitle = attr.Val
				}
			}

			if isOgImage(node.Attr) {
				for _, attr := range node.Attr {
					ogImage = attr.Val
				}
			}
		}

		for fc := node.FirstChild; fc != nil; fc = fc.NextSibling {
			f(fc)
		}
	}

	f(doc)

	return &Page{
		Title:       title,
		Description: desc,
		OgImage:     ogImage,
		OgTitle:     ogTitle,
	}, nil
}
