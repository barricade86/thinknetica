package main

import (
	"fmt"
	"sync"
	"thinknetica/pingpong/pkg/storage"
)

func main() {
	ppChan := make(chan string, 1)
	var wg sync.WaitGroup
	matchStorage := storage.New()
	firstPlayerName := "Freddy"
	secondPlayerName := "Jason"
	matchStorage.CreatePlayer(firstPlayerName)
	matchStorage.CreatePlayer(secondPlayerName)
	wg.Add(1)
	go play(ppChan, matchStorage, firstPlayerName, &wg)
	wg.Add(1)
	go play(ppChan, matchStorage, secondPlayerName, &wg)
	ppChan <- "begin"
	wg.Wait()
	wg.Add(1)
	go func(wg *sync.WaitGroup, matchStorage *storage.Match) {
		defer wg.Done()
		message := <-ppChan
		if message == "stop" {
			fmt.Printf("%v", matchStorage.GetTotalScore())
		}
	}(&wg, matchStorage)
	wg.Wait()
}

func play(ppChan chan string, matchStorage *storage.Match, playerName string, wg *sync.WaitGroup) {
	defer wg.Done()
	var message string
	for {
		message = <-ppChan
		fmt.Printf("Move from %s %s \n", message, playerName)
		player, err := matchStorage.GetPlayerByName(playerName)
		if err != nil {
			fmt.Printf("Error getting player by name %s\n", err)
			continue
		}

		if player.Score >= 11 {
			ppChan <- "stop"
			break
		}

		if message == "begin" || message == "pong" {
			matchStorage.AddPoint(player.Name, 1)
			ppChan <- "ping"
		}

		if message == "ping" {
			matchStorage.AddPoint(player.Name, 1)
			ppChan <- "pong"
		}
	}
}
