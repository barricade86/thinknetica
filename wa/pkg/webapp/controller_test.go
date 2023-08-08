package webapp

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"thinknetica/wa/pkg/crawler"
	"thinknetica/wa/pkg/storage"

	"github.com/gorilla/mux"
)

func TestAdd(t *testing.T) {
	expected := `OK`
	req := httptest.NewRequest(http.MethodPost, "/add", strings.NewReader(`{"title": "Transformers", "body": "One shall rise another shall fall", "url": "http://transformers.com"}`))
	w := httptest.NewRecorder()
	c := NewController(storage.NewInMemoryStorage())
	c.Add(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if string(data) != expected {
		t.Errorf("Expected %s but got %v", expected, string(data))
	}
}

func TestRemove(t *testing.T) {
	expected := `OK`
	storage := storage.NewInMemoryStorage()
	storage.Add(crawler.Document{ID: 1, Title: "Transformers", Body: "One shall rise another shall fall", URL: "http://transformers.com"})
	req := httptest.NewRequest(http.MethodPost, "/delete/1", nil)
	w := httptest.NewRecorder()
	c := NewController(storage)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	c.Remove(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if string(data) != expected {
		t.Errorf("Expected %s but got %v", expected, string(data))
	}
}

func TestUpdateById(t *testing.T) {
	expected := `OK`
	storage := storage.NewInMemoryStorage()
	storage.Add(crawler.Document{ID: 1, Title: "Transformers:Revenge of the fallen", Body: "One shall rise another shall fall", URL: "http://transformers.com"})
	req := httptest.NewRequest(http.MethodPost, "/update", strings.NewReader(`{"title": "Transformers.Revenge of the fallen", "body": "One shall rise another shall fall", "url": "http://transformers.com"}`))
	w := httptest.NewRecorder()
	c := NewController(storage)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	c.UpdateById(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if string(data) != expected {
		t.Errorf("Expected %s but got %v", expected, string(data))
	}
}
