package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		fmt.Printf("Server connection error:%s", err)
		return
	}
	defer conn.Close()
	var message []byte
	var needle string
	reader := bufio.NewReader(os.Stdin)
	serverReader := bufio.NewReader(conn)
	for {
		fmt.Print("Input search string ")
		needle, err = reader.ReadString('\n')
		fmt.Println("You entered ", needle)
		switch err {
		case io.EOF:
			fmt.Printf("client closed the connection")
			break
		case nil:
			n, err := conn.Write([]byte(needle + "\n"))
			if err != nil {
				fmt.Printf("Error sending request:%s", err)
			}
			fmt.Printf("%d bytes was sent to server\n", n)
			break
		default:
			fmt.Printf("server error: %v\n", err)
			break
		}

		message, err = serverReader.ReadBytes('\n')
		if len(message) > 0 {
			fmt.Printf("Search results:%s\n", message)
		}
	}
}
