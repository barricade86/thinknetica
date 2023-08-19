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
	fmt.Printf("%v", matchStorage.GetTotalScore())
}

func play(ppChan chan string, matchStorage *storage.Match, playerName string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		message, ok := <-ppChan
		if !ok {
			break
		}

		fmt.Printf("Move from %s %s \n", playerName, message)
		if message == "stop" {
			close(ppChan)
			break
		}

		player, err := matchStorage.GetPlayerByName(playerName)
		if err != nil {
			fmt.Printf("Error getting player by name %s\n", err)
			continue
		}
		fmt.Printf("Player score  %d\n", player.Score)
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
