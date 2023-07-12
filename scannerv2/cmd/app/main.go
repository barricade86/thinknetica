package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sort"
	"thinknetica/scannerv1/pkg/crawler"
	"thinknetica/scannerv1/pkg/crawler/spider"
	"thinknetica/scannerv2/pkg/index"
	"time"
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
	for _, url := range resources {
		result, err := spider.Scan(url, depth)
		if err != nil {
			fmt.Printf("Error due to scanning docs in %s resourse: %s", url, err)
			continue
		}

		scanResults = append(scanResults, result...)
	}

	source := rand.NewSource(time.Now().UnixNano())
	randSource := rand.New(source)
	for idx, _ := range scanResults {
		scanResults[idx].ID = randSource.Int()
	}

	sort.Slice(scanResults, func(i, j int) bool {
		return scanResults[i].ID < scanResults[j].ID
	})
	reverseIndexService := index.NewReverseIndex()
	reverseIndex := reverseIndexService.Build(scanResults)
	docIds, ok := reverseIndex[needle]
	if !ok {
		fmt.Printf("No data found by key %s", needle)
		return
	}

	for _, docId := range docIds {
		idx := sort.Search(len(scanResults), func(i int) bool { return scanResults[i].ID >= docId })
		if idx < len(scanResults) && scanResults[idx].ID == docId {
			fmt.Printf("%s %s\n", scanResults[idx].Title, scanResults[idx].Body)
		}
	}
}
