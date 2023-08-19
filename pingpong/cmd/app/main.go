package main

import (
	"fmt"
	"math/rand"
	"sync"
	"thinknetica/pingpong/pkg/storage"
)

func main() {
	ppChan := make(chan string, 1)
	var wg sync.WaitGroup
	matchStorage := storage.New()
	var players = map[uint]string{0: "Freddy", 1: "Jason"}
	randomIdx := rand.Intn(2)
	var first, second uint
	if randomIdx == 1 {
		first, second = 0, 1
	} else {
		first, second = 1, 0
	}

	firstPlayerName := players[first]
	secondPlayerName := players[second]
	matchStorage.CreatePlayer(firstPlayerName)
	matchStorage.CreatePlayer(secondPlayerName)
	wg.Add(1)
	go play(ppChan, matchStorage, firstPlayerName, &wg)
	wg.Add(1)
	go play(ppChan, matchStorage, secondPlayerName, &wg)
	ppChan <- "begin"
	wg.Wait()
	fmt.Printf("%v", matchStorage.GetTotalScore())
}

func play(ppChan chan string, matchStorage *storage.Match, playerName string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		message, ok := <-ppChan
		if !ok {
			break
		}

		if matchStorage.ContainsValue(11) {
			break
		}

		num := rand.Int()
		if num < 20 {
			break
		}

		fmt.Printf("Move from %s %s \n", playerName, message)
		if message == "stop" {
			close(ppChan)
			break
		}

		if message == "begin" || message == "pong" {
			matchStorage.AddPoint(playerName, 1)
			ppChan <- "ping"
		}

		if message == "ping" {
			matchStorage.AddPoint(playerName, 1)
			ppChan <- "pong"
		}
	}

	ppChan <- "stop"
}
