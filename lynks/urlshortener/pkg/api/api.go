package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"lynks/urlshortener/pkg/response"
	"lynks/urlshortener/pkg/storage"
	"net/http"
)

type API struct {
	pgStorage *storage.ShortLinks
}

func New(storage *storage.ShortLinks) *API {
	return &API{pgStorage: storage}
}

func (a *API) CreateShortLink(w http.ResponseWriter, r *http.Request) {
	requestBodyReader, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Incorrect or malformed request data`))
		return
	}

	originalLinkMap := make(map[string]string)
	err = json.Unmarshal(requestBodyReader, &originalLinkMap)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error unmarshalling data:%s", err)))
		return
	}

	link := originalLinkMap["original"]
	var shortLink string
	shortLink, err = a.pgStorage.Add(link)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Generating link error:%s", err)))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{"originalLink":%s,"shortLink":%s}`, link, shortLink)))
	return
}

func (a *API) RedirectToOriginal(w http.ResponseWriter, r *http.Request) {
	requestBodyReader, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Incorrect or malformed request data`))
		return
	}

	link := &response.Link{}
	err = json.Unmarshal(requestBodyReader, link)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error unmarshalling data:%s", err)))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{"originalLink":%s,"shortLink":%s}`, link.Original, link.Short)))
	return
}
