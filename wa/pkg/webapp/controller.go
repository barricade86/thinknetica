package webapp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"thinknetica/wa/pkg/crawler"
	"thinknetica/wa/pkg/storage"
	"thinknetica/webapplication/pkg/webapp/response"
	"time"

	"github.com/gorilla/mux"
)

type Controller struct {
	storage *storage.InMemoryStorage
}

func NewController(storage *storage.InMemoryStorage) *Controller {
	return &Controller{storage: storage}
}

func (c *Controller) Add(w http.ResponseWriter, r *http.Request) {
	docRequestData := &response.DocData{}
	requestBodyReader, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Incorrect or malformed request body`))
		return
	}

	err = json.Unmarshal(requestBodyReader, docRequestData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(fmt.Sprintf("error unmarshal data:%s. Request body %s", err, requestBodyReader))
		return
	}

	source := rand.NewSource(time.Now().UnixNano())
	randSource := rand.New(source)
	id := randSource.Int()
	c.storage.Add(crawler.Document{ID: id, Title: docRequestData.Title, URL: docRequestData.URL, Body: docRequestData.Body})
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`OK`))
}

func (c *Controller) Remove(w http.ResponseWriter, r *http.Request) {
	queryParams := mux.Vars(r)
	_, ok := queryParams["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Missing id parameter`))
		return
	}

	docID, err := strconv.Atoi(queryParams["id"])
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Incorrect or malformed id parameter`))
		return
	}

	err = c.storage.Delete(docID)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("Document with %d not found", docID)))
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("record with id %d not exists", docID)))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`OK`))
}

func (c *Controller) FindByQueryText(w http.ResponseWriter, r *http.Request) {
	queryParams := mux.Vars(r)
	queryText, ok := queryParams["queryText"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Missing queryText parameter`))
		return
	}

	documents := c.storage.FindByQueryText(queryText)
	if len(documents) == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("No data found by query %s", queryText)))
		return
	}

	docs := make([]*response.DocData, 0)
	for _, val := range documents {
		docs = append(docs, &response.DocData{Title: val.Title, Body: val.Body, URL: val.URL})
	}

	result, err := json.Marshal(docs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Marshalling error %s", err)))
		return
	}

	_, err = w.Write(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(result)
}

func (c *Controller) UpdateById(w http.ResponseWriter, r *http.Request) {
	queryParams := mux.Vars(r)
	_, ok := queryParams["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Missing id parameter`))
		return
	}

	docID, err := strconv.Atoi(queryParams["id"])
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Incorrect or malformed id parameter`))
		return
	}

	docData := crawler.Document{}
	requestBodyReader, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Incorrect or malformed request body`))
		return
	}

	err = json.Unmarshal(requestBodyReader, &docData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = c.storage.UpdateById(docID, docData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Update error %s", err)))
		return
	}

	w.Write([]byte(`OK`))
}
