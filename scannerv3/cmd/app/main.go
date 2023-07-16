package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"thinknetica/scannerv3/pkg/crawler"
	"thinknetica/scannerv3/pkg/crawler/spider"
)

const depth = 2
const pathToFile = "../../data.txt"

var needle string

var resources = []string{"https://golang-org.appspot.com/", "https://go.dev/"}

func initFlags() {
	flag.StringVar(&needle, "needle", "", "Option sets search key")
	flag.Parse()
}

func main() {
	initFlags()
	if needle == "" {
		fmt.Println("Option needle is empty")
		return
	}

	_, err := os.Stat(pathToFile)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			initFileStorage()
		} else {
			panic(err)
		}
	}

	searchResults, err := search(needle, pathToFile)
	if len(searchResults) == 0 {
		fmt.Println("No data found")
		return
	}

	for _, value := range searchResults {
		fmt.Println(value)
	}

	fmt.Printf("Total results found:%d", len(searchResults))
}

func search(needle string, pathToFile string) ([]string, error) {
	file, err := os.Open(pathToFile)
	if err != nil {
		return []string{}, nil
	}

	var links []string
	scanner := bufio.NewReader(file)
	for {
		buf, _, err := scanner.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return []string{}, err
			}
		}

		line := fmt.Sprintf("%s", buf)
		slicedResult := strings.Split(line, "||")
		if strings.Contains(slicedResult[0], needle) {
			links = append(links, slicedResult[1])
		}
	}

	return links, err
}

// initFileStorage
func initFileStorage() {
	file, err := os.OpenFile(pathToFile, os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		panic(err)
	}

	defer file.Close()
	if err := os.Truncate(pathToFile, 0); err != nil {
		panic(fmt.Errorf("Failed to truncate: %v", err))
	}

	spider := spider.New()
	scanResults := make([]crawler.Document, 0)
	for _, val := range resources {
		result, err := spider.Scan(val, depth)
		if err != nil {
			fmt.Errorf("Error due to scanning docs in %s resourse: %s", val, err)
			continue
		}

		scanResults = append(scanResults, result...)
	}

	for _, result := range scanResults {
		resultStr := ""
		if result.Title != "" {
			resultStr = resultStr + result.Title
		}

		if result.Body != "" {
			resultStr = resultStr + "||" + result.Body
		}

		if result.URL != "" {
			resultStr = resultStr + "||" + result.URL
		}

		_, err := file.Write([]byte(resultStr + "\n"))
		if err != nil {
			panic(fmt.Errorf("Error due to write data in file: %s", err))
		}
	}
}
