package main

import (
	"net/http"
	"thinknetica/wa/pkg/storage"
	"thinknetica/wa/pkg/webapp"

	"github.com/gorilla/mux"
)

const depth = 2

var needle string

func main() {
	storage := storage.NewInMemoryStorage()
	api := webapp.New(storage)
	router := mux.NewRouter()
	router.HandleFunc("/add", api.Add)
	router.HandleFunc("/delete/{id}", api.Remove)
	router.HandleFunc("/show/{queryText}", api.FindByQueryText)
	router.HandleFunc("/update/{id}", api.UpdateById)
	http.ListenAndServe("localhost:8081", router)
}
