package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sort"
	"thinknetica/wa/pkg/crawler"
	"thinknetica/wa/pkg/crawler/spider"
	"thinknetica/wa/pkg/index"
	"thinknetica/wa/pkg/storage"
	"thinknetica/wa/pkg/webapp"
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
	_ = reverseIndexService.Build(scanResults)
	storage := storage.NewInMemoryStorage()
	controller := webapp.NewController(storage)
	router := mux.NewRouter()
	router.HandleFunc("/add", controller.Add)
	router.HandleFunc("/delete/{id}", controller.Remove)
	router.HandleFunc("/show/{queryText}", controller.FindByQueryText)
	router.HandleFunc("/update/{id}", controller.UpdateById)
	log.Fatalf("%s", http.ListenAndServe("localhost:8081", router))
}
