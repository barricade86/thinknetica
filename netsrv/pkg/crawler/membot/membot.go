package membot

import (
	"thinknetica/searcher/pkg/crawler"
)

// Service - имитация служба поискового робота.
type Service struct{}

// New - констрктор имитации службы поискового робота.
func New() *Service {
	s := Service{}
	return &s
}

// Scan возвращает заранее подготовленный набор данных
func (s *Service) Scan(url string, depth int) ([]crawler.Document, error) {

	data := []crawler.Document{
		{
			ID:    0,
			URL:   "https://go.dev/",
			Title: "go-dev",
		},
		{
			ID:    1,
			URL:   "https://golang-org.appspot.com/",
			Title: "golang org",
		},
	}

	return data, nil
}
