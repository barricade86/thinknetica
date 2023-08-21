package netsrv

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"thinknetica/cs/pkg/crawler"
	"time"
)

type Server struct {
	netListener net.Listener
	scanResults []crawler.Document
}

func NewServer(netListener net.Listener, scanResults []crawler.Document) *Server {
	return &Server{netListener: netListener, scanResults: scanResults}
}

func (s *Server) ListenAndServe() {
	defer s.netListener.Close()
	fmt.Println("Server is ready to accept client connections")
	for {
		conn, err := s.netListener.Accept()
		if err != nil {
			fmt.Printf("Accept connection error:%s", err)
			continue
		}

		go handleClientRequest(conn, s.scanResults)
	}
}

func handleClientRequest(connection net.Conn, scanResults []crawler.Document) {
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
