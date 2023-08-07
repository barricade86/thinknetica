package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"thinknetica/cs/pkg/crawler"
	"thinknetica/cs/pkg/crawler/spider"
	"thinknetica/cs/pkg/netsrv"
	"time"
)

const depth = 2

type ConnectionHandler func(net.Conn, []crawler.Document)

func main() {
	resources := []string{"https://golang-org.appspot.com/", "https://go.dev/"}
	spider := spider.New()
	var scanResults []crawler.Document
	fmt.Println("Start scanning resorces")
	for _, url := range resources {
		result, err := spider.Scan(url, depth)
		if err != nil {
			fmt.Printf("Error due to scanning docs in %s resourse: %s", url, err)
			continue
		}

		scanResults = append(scanResults, result...)
	}

	fmt.Println("Scanning is finished")
	handler := func(connection net.Conn, scanResults []crawler.Document) {
		defer connection.Close()
		var err error
		var needle string
		reader := bufio.NewReader(connection)
		searchFunc := func(needle string, scanResults []crawler.Document) []string {
			var links []string
			for _, value := range scanResults {
				if strings.Contains(value.Title, needle) || strings.Contains(value.Body, needle) {
					links = append(links, fmt.Sprintf("%s %s\n", value.Title, value.Body))
				}
			}

			return links
		}

		for {
			needle, err = reader.ReadString('\n')
			if err != nil {
				fmt.Printf("Error reading client request:%s \n", err)
				return
			}

			needle = strings.TrimSpace(needle)
			searchResult := searchFunc(needle, scanResults)
			if len(searchResult) == 0 {
				result := fmt.Sprintf("No data found by : %s", needle)
				_, err := connection.Write([]byte(result + "\n"))
				if err != nil {
					fmt.Printf("error writing response to client:%s", err)
				}

				continue
			}

			result := strings.Join(searchResult, ",")
			_, err = connection.Write([]byte(result + "\n"))

			if err != nil {
				fmt.Printf("error writing response to client:%s", err)
			}

			connection.SetDeadline(time.Now().Add(time.Second * 60))
		}
	}

	fmt.Println("Starting tcp server")
	listener, err := net.Listen("tcp", "0.0.0.0:8000")
	if err != nil {
		fmt.Printf("Listener error:%s", err)
		return
	}

	netListener := netsrv.NewServer(listener, scanResults, handler)
	netListener.ListenAndServe()
	defer listener.Close()
}
