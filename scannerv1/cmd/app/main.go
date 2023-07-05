package main

import (
	"flag"
	"fmt"
	"thinkneticacourse/scannerv1/pkg/crawler/spider"
)

var needle string
var depth int

func init() {
	flag.StringVar(&needle, "needle", "", "Option sets search key")
	flag.IntVar(&depth, "depth", 0, "Option sets depth of search")
	flag.Parse()
}

func main() {
	if needle == "" {
		fmt.Println("Option needle is empty")
		return
	}

	if depth == 0 {
		fmt.Println("Option depth is empty")
		return
	}

	spider := spider.New()
	resultFirst, err := spider.Scan("https://golang-org.appspot.com/", depth)
	if err != nil {
		fmt.Printf("Error due to scanning docs in golang-org resourse", err)
		return
	}

	for _, value := range resultFirst {
		fmt.Println(value.URL)
	}

	resultSecond, err := spider.Scan("https://go.dev/", depth)
	if err != nil {
		fmt.Printf("Error due to scanning docs in golang-org resourse", err)
		return
	}

	for _, value := range resultSecond {
		fmt.Println(value.URL)
	}

	fmt.Printf("Total results found:%d", len(resultFirst)+len(resultSecond))
}
