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
	fmt.Println("Starting tcp server")
	listener, err := net.Listen("tcp", "0.0.0.0:8000")
	if err != nil {
		fmt.Printf("Listener error:%s", err)
		return
	}

	defer listener.Close()
	fmt.Println("Server is ready to accept client connections")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Accept connection error:%s", err)
			continue
		}

		go handler(conn, scanResults)
	}
}

func handlertwo(connection net.Conn) {
	reader := bufio.NewReader(connection)
	for {
		needle, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading client request:%s \n", err)
			return
		}

		n, err := connection.Write([]byte(fmt.Sprintf("got it %s\n", needle)))
		if err != nil {
			fmt.Printf("error writing response to client:%s", err)
		}

		fmt.Printf("%d bytes was sent", n)
	}
}

func handler(connection net.Conn, scanResults []crawler.Document) {
	defer connection.Close()
	var err error
	var needle string
	reader := bufio.NewReader(connection)
	for {
		needle, err = reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading client request:%s \n", err)
			return
		}

		needle = strings.TrimSpace(needle)
		searchResult := search(needle, scanResults)
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

func search(needle string, scanResults []crawler.Document) []string {
	var links []string
	for _, value := range scanResults {
		if strings.Contains(value.Title, needle) || strings.Contains(value.Body, needle) {
			links = append(links, fmt.Sprintf("%s %s\n", value.Title, value.Body))
		}
	}

	return links
}
