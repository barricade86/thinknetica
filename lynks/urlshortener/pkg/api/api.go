package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"lynks/common"
	"lynks/urlshortener/pkg/response"
	"net/http"

	"github.com/rs/zerolog/log"
)

type API struct {
	pgStorage common.Storage
}

func New(storage common.Storage) *API {
	return &API{pgStorage: storage}
}

func (a *API) CreateShortLink(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(r.Method + ` Method not allowed`))
		return
	}

	log.Info().Msg("Reading request body")
	requestBodyReader, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error().Any("request_body", r.Body).AnErr("read_request_body_err", err).Msg("Incorrect or malformed request data")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Incorrect or malformed request data`))
		return
	}

	log.Info().Msg("Request body read")
	originalLinkMap := make(map[string]string)
	log.Info().Msg("Unmarshalling request body")
	err = json.Unmarshal(requestBodyReader, &originalLinkMap)
	if err != nil {
		log.Error().AnErr("unmarshal_request_body_err", err).Msg("Error unmarshalling data")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid json`))
		return
	}

	log.Info().Msg("Request body unmarshalled")
	link := originalLinkMap["original"]
	var shortLink string
	log.Info().Msg("Saving link to storage")
	shortLink, err = a.pgStorage.Add(link)
	if err != nil {
		log.Error().AnErr("generate_link_err", err).Msg("Generating link error")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`Generating link error`))
		return
	}

	log.Info().Msg("Link saved to storage")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{"originalLink":%s,"shortLink":%s}`, link, shortLink)))
	return
}

func (a *API) RedirectToOriginal(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(r.Method + ` Method not allowed`))
		return
	}

	log.Info().Msg("Start reading request body")
	requestBodyReader, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error().Any("request_body", r.Body).AnErr("read_request_body_err", err).Msg("Incorrect or malformed request data")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Incorrect or malformed request data`))
		return
	}

	log.Info().Msg("Reading request body finished")
	link := &response.Link{}
	log.Info().Msg("Unmarshalling request body")
	err = json.Unmarshal(requestBodyReader, link)
	if err != nil {
		log.Error().AnErr("unmarshal_request_body_err", err).Msg("Error unmarshalling data")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid json`))
		return
	}

	log.Info().Msg("Request body unmarshalled")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{"originalLink":%s,"shortLink":%s}`, link.Original, link.Short)))
	return
}
