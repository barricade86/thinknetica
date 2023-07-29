package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"thinknetica/netsrv/pkg/crawler"
	"thinknetica/netsrv/pkg/crawler/spider"
	"time"
)

const depth = 2

func main() {
	resources := []string{"https://golang-org.appspot.com/", "https://go.dev/"}
	spider := spider.New()
	var scanResults []crawler.Document
	for _, url := range resources {
		result, err := spider.Scan(url, depth)
		if err != nil {
			fmt.Printf("Error due to scanning docs in %s resourse: %s", url, err)
			continue
		}

		scanResults = append(scanResults, result...)
	}

	listener, err := net.Listen("tcp4", "0.0.0.0:8000")
	if err != nil {
		fmt.Printf("Listener error:%s", err)
		return
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Accept connection error:%s", err)
			return
		}

		go handler(conn, scanResults)
	}

}

func handler(connection net.Conn, scanResults []crawler.Document) {
	var err error
	var needle string
	reader := bufio.NewReader(connection)
	needle, err = reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading client request:%s \n", err)
		return
	}

	connection.SetDeadline(time.Now().Add(time.Second * 60))
	needle = strings.TrimSpace(needle)
	fmt.Printf("Client request: %s \n", needle)
	searchResult := search(needle, scanResults)
	if len(searchResult) == 0 {
		result := fmt.Sprintf("No data found by : %s", needle)
		_, err := connection.Write([]byte(result))
		if err != nil {
			fmt.Printf("error writing response to client:%s", err)
		}
		connection.Close()
		return
	}

	result := strings.Join(searchResult, ",")
	_, err = connection.Write([]byte(result))
	connection.Close()
	if err != nil {
		fmt.Printf("error writing response to client:%s", err)
	}
}

func search(needle string, scanResults []crawler.Document) []string {
	var links []string
	for _, value := range scanResults {
		if strings.Contains(value.Title, needle) || strings.Contains(value.Body, needle) {
			links = append(links, fmt.Sprintf("%s %s\n", value.Title, value.Body))
		}
	}

	return links
}
