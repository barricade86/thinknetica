package main

import (
	"fmt"
	"thinknetica/pingpong/pkg/storage"
)

func main() {
	//var wg sync.WaitGroup
	pingCh := make(chan string)
	pongCh := make(chan string)
	playerStorage := storage.New()
	playerStorage.Create("Freddy")
	playerStorage.Create("Jason")
	go ping(pingCh, pongCh, playerStorage)
	go pong(pongCh, pingCh, playerStorage)
	pingCh <- "ping"
	/*wg.Add(1)
	resultCh := pong(pingCh)
	pingCh <- "ping"
	go func(pongCh chan string, wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Println("pongResult = ", <-pongCh)
	}(resultCh, &wg)
	wg.Wait()*/
}

func ping(chanIn chan string, chanOut chan string, playerStorage *storage.Player) {
	for {
		result := <-chanIn
		fmt.Println("chanIn result=", result)
		chanOut <- "pong"
	}
}

func pong(chanIn chan string, chanOut chan string, playerStorage *storage.Player) {
	for {
		result := <-chanIn
		fmt.Println("chanOut result=", result)
		chanOut <- "ping"
	}
}
