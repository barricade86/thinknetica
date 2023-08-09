package webapp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"thinknetica/webapplication/pkg/crawler"
	"thinknetica/webapplication/pkg/webapp/response"
)

type Controller struct {
	index       map[string][]int
	scanResults []crawler.Document
}

func NewController(index map[string][]int, scanResults []crawler.Document) *Controller {
	return &Controller{index: index, scanResults: scanResults}
}

func (c *Controller) ShowIndexData(w http.ResponseWriter, r *http.Request) {
	var indexData []*response.IndexData
	for key, val := range c.index {
		indexData = append(indexData, &response.IndexData{Token: key, PositionList: fmt.Sprintf("%v", val)})
	}

	result, err := json.Marshal(indexData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func (c *Controller) ShowDocData(w http.ResponseWriter, r *http.Request) {
	var indexData []*response.DocData
	for _, val := range c.scanResults {
		indexData = append(indexData, &response.DocData{Title: val.Title, Body: val.Body, URL: val.URL})
	}

	result, err := json.Marshal(indexData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
