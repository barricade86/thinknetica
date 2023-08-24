package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	ppChan := make(chan string, 1)
	var wg sync.WaitGroup
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
	go play(ppChan, tournamentTable, firstPlayerName, &wg)
	wg.Add(1)
	go play(ppChan, tournamentTable, secondPlayerName, &wg)
	ppChan <- "begin"
	wg.Wait()
	fmt.Printf("%v", tournamentTable)
}

func play(ppChan chan string, tournamentTable map[string]uint, playerName string, wg *sync.WaitGroup) {
	defer wg.Done()
	var nextMove string
	for move := range ppChan {
		if ContainsValue(tournamentTable, 11) {
			break
		}

		num := rand.Intn(20)
		if num == 20 {
			break
		}

		fmt.Printf("Move from %s %s \n", playerName, move)
		switch move {
		case "stop":
			close(ppChan)
			break
		case "begin", "pong":
			nextMove = "ping"
			break
		case "ping":
			nextMove = "pong"
		default:
			continue
		}

		if nextMove != "" {
			tournamentTable[playerName]++
			ppChan <- nextMove
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
