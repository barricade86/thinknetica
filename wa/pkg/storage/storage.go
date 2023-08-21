package storage

import (
	"fmt"
	"strings"
	"sync"
	"thinknetica/wa/pkg/crawler"
)

type InMemoryStorage struct {
	storage map[int]crawler.Document
	mu      sync.RWMutex
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{storage: make(map[int]crawler.Document, 0)}
}

func (ims *InMemoryStorage) Add(item crawler.Document) {
	ims.mu.Lock()
	ims.storage[item.ID] = item
	defer ims.mu.Unlock()
}

func (ims *InMemoryStorage) FindByQueryText(query string) []crawler.Document {
	var result []crawler.Document
	for _, val := range ims.storage {
		if strings.Contains(val.Title, query) || strings.Contains(val.Body, query) {
			result = append(result, val)
		}
	}

	return result
}

func (ims *InMemoryStorage) Delete(id int) error {
	_, ok := ims.storage[id]
	if !ok {
		return fmt.Errorf("record with id %d not exists", id)
	}

	ims.mu.Lock()
	delete(ims.storage, id)
	defer ims.mu.Unlock()
	return nil
}

func (ims *InMemoryStorage) UpdateById(id int, item crawler.Document) error {
	_, ok := ims.storage[id]
	if !ok {
		return fmt.Errorf("record with id %d not exists", id)
	}

	ims.mu.Lock()
	ims.storage[item.ID] = item
	defer ims.mu.Unlock()

	return nil
}
