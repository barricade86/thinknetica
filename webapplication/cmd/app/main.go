package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sort"
	"thinknetica/webapplication/pkg/crawler"
	"thinknetica/webapplication/pkg/crawler/spider"
	"thinknetica/webapplication/pkg/index"
	"thinknetica/webapplication/pkg/webapp"
	"time"

	"github.com/gorilla/mux"
)

const depth = 2

var needle string

func main() {
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
	controller := webapp.NewController(reverseIndex, scanResults)
	router := mux.NewRouter()
	router.HandleFunc("/index", controller.ShowIndexData)
	router.HandleFunc("/docs", controller.ShowDocData)
	log.Fatalf("%s", http.ListenAndServe("localhost:8081", router))
}
