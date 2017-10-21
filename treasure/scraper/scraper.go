package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

type Page struct {
	Title       string
	Description string
}

func isDescription(attrs []html.Attribute) bool {
	for _, attr := range attrs {
		if attr.Key == "name" && attr.Val == "description" {
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
		}

		for fc := node.FirstChild; fc != nil; fc = fc.NextSibling {
			f(fc)
		}
	}

	f(doc)

	return &Page{Title: title, Description: desc}, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "URL must specified")
		os.Exit(1)
	}

	p, err := Get(os.Args[1])
	if err != nil && err != io.EOF {
		panic(err)
	}
	fmt.Printf("%#v", p)
}
