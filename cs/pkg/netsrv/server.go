package netsrv

import (
	"fmt"
	"net"
	"thinknetica/cs/pkg/crawler"
)

type Server struct {
	netListener net.Listener
	handler     func(net.Conn, []crawler.Document)
	scanResults []crawler.Document
}

func NewServer(netListener net.Listener, scanResults []crawler.Document, handler func(net.Conn, []crawler.Document)) *Server {
	return &Server{netListener: netListener, scanResults: scanResults, handler: handler}
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

		go s.handler(conn, s.scanResults)
	}
}
