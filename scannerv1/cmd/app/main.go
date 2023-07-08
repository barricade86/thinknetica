package main

import (
	"flag"
	"fmt"
	"strings"
	"thinknetica/scannerv1/pkg/crawler"
	"thinknetica/scannerv1/pkg/crawler/spider"
)

const depth = 2

var needle string

func initFlags() {
	flag.StringVar(&needle, "needle", "", "Option sets search key")
	flag.Parse()
}

func main() {
	initFlags()
	if needle == "" {
		fmt.Println("Option needle is empty")
		return
	}

	resources := []string{"https://golang-org.appspot.com/", "https://go.dev/"}
	spider := spider.New()
	var scanResults []crawler.Document
	for _, val := range resources {
		result, err := spider.Scan(val, depth)
		if err != nil {
			fmt.Printf("Error due to scanning docs in %s resourse: %s", val, err)
			continue
		}

		scanResults = append(scanResults, result...)
	}

	searchResults := search(needle, scanResults)
	if len(searchResults) == 0 {
		fmt.Println("No data found")
		return
	}

	for _, value := range searchResults {
		fmt.Println(value)
	}

	fmt.Printf("Total results found:%d", len(searchResults))
}

func search(needle string, scanResults []crawler.Document) []string {
	var links []string
	for _, value := range scanResults {
		if strings.Contains(value.Title, needle) || strings.Contains(value.Body, needle) {
			links = append(links, value.URL)
		}
	}

	return links
}
