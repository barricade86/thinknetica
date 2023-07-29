package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp4", "localhost:8000")
	if err != nil {
		fmt.Printf("Server connection error:%s", err)
		return
	}
	defer conn.Close()
	var message []byte
	var needle string
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Input search string ")
		needle, _ = reader.ReadString('\n')
		fmt.Println("You entered ", needle)
		conn.Write([]byte(needle))
		message, err = io.ReadAll(conn)
		if err != nil {
			fmt.Printf("Server connection read data error:%s", err)
			return
		}

		if len(message) > 0 {
			fmt.Printf("Search results:%s", message)
		}
	}
}
