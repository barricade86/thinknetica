package index

import (
	"strings"
	"thinknetica/webapplication/pkg/crawler"
)

type ReverseIndex struct {
	storage map[string][]int
}

func NewReverseIndex() *ReverseIndex {
	return &ReverseIndex{storage: make(map[string][]int)}
}

func (r *ReverseIndex) Build(data []crawler.Document) map[string][]int {
	//tokens:=r.tokenize(data)
	var tokens []string
	result := make(map[string][]int)
	for _, value := range data {
		tokens = strings.Split(value.Title, " ")
		for _, token := range tokens {
			if token == "-" || token == "&" || token == "." {
				continue
			}

			_, ok := result[token]
			if !ok {
				result[token] = []int{value.ID}
			} else {
				result[token] = append(result[token], value.ID)
			}
		}
	}

	return result
}
