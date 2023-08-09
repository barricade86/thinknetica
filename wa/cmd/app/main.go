package main

import (
	"log"
	"net/http"
	"thinknetica/wa/pkg/storage"
	"thinknetica/wa/pkg/webapp"

	"github.com/gorilla/mux"
)

const depth = 2

var needle string

func main() {
	storage := storage.NewInMemoryStorage()
	controller := webapp.NewController(storage)
	router := mux.NewRouter()
	router.HandleFunc("/add", controller.Add)
	router.HandleFunc("/delete/{id}", controller.Remove)
	router.HandleFunc("/show/{queryText}", controller.FindByQueryText)
	router.HandleFunc("/update/{id}", controller.UpdateById)
	log.Fatalf("%s", http.ListenAndServe("localhost:8081", router))
}
