package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	ppChan := make(chan string, 1)
	var wg sync.WaitGroup
	var mu sync.RWMutex
	var players = map[uint]string{0: "Freddy", 1: "Jason"}
	var tournamentTable = map[string]uint{"Freddy": 0, "Jason": 0}
	randomIdx := rand.Intn(2)
	var first, second uint
	if randomIdx == 1 {
		first, second = 0, 1
	} else {
		first, second = 1, 0
	}

	firstPlayerName := players[first]
	secondPlayerName := players[second]
	wg.Add(1)
	go play(ppChan, tournamentTable, firstPlayerName, &wg, mu)
	wg.Add(1)
	go play(ppChan, tournamentTable, secondPlayerName, &wg, mu)
	ppChan <- "begin"
	wg.Wait()
	fmt.Printf("%v", tournamentTable)
}

func play(ppChan chan string, tournamentTable map[string]uint, playerName string, wg *sync.WaitGroup, mu sync.RWMutex) {
	defer wg.Done()
	for {
		message, ok := <-ppChan
		if !ok {
			break
		}

		if ContainsValue(tournamentTable, 11) {
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
			mu.Lock()
			tournamentTable[playerName]++
			ppChan <- "ping"
			mu.Unlock()
		}

		if message == "ping" {
			mu.Lock()
			tournamentTable[playerName]++
			ppChan <- "pong"
			mu.Unlock()
		}
	}

	ppChan <- "stop"
}

func ContainsValue(tournamentTable map[string]uint, value uint) bool {
	for _, x := range tournamentTable {
		if x == value {
			return true
		}
	}
	return false
}
