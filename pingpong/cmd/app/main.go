package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	pingCh := make(chan string)
	wg.Add(1)
	resultCh := pong(pingCh)
	pingCh <- "ping"
	go func(pongCh chan string, wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Println("pongResult = ", <-pongCh)
	}(resultCh, &wg)
	wg.Wait()
}

func pong(ch chan string) chan string {
	pongCh := make(chan string)
	go func(pingCh chan string) {
		result := <-pingCh
		fmt.Println("result=", result)
		if result == "ping" {
			pongCh <- "pong"
		}
	}(ch)

	return pongCh
}
