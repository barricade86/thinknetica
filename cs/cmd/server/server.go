package main

import (
	"fmt"
	"net"
	"thinknetica/cs/pkg/crawler"
	"thinknetica/cs/pkg/crawler/spider"
	"thinknetica/cs/pkg/netsrv"
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
	fmt.Println("Starting tcp server")
	listener, err := net.Listen("tcp", "0.0.0.0:8000")
	if err != nil {
		fmt.Printf("Listener error:%s", err)
		return
	}

	netListener := netsrv.NewServer(listener, scanResults)
	netListener.ListenAndServe()
	defer listener.Close()
}
