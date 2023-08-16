package main

import (
	"fmt"
	"thinknetica/pingpong/pkg/model"
	"thinknetica/pingpong/pkg/storage"
)

func main() {
	pingCh := make(chan string)
	pongCh := make(chan string)
	playerStorage := storage.New()
	freddy := playerStorage.Create("Freddy")
	jason := playerStorage.Create("Jason")
	go play(pingCh, pongCh, freddy)
	go play(pongCh, pingCh, jason)
	pingCh <- "begin"
	for {
		select {
		case message := <-pongCh:
			fmt.Println("message=", <-pongCh)
			if message == "stop" {
				return
			}
		case message := <-pingCh:
			fmt.Println("message=", <-pingCh)
			if message == "stop" {
				return
			}
		default:
			continue
		}
	}
}

func play(chanIn chan string, chanOut chan string, player *model.Player) {
	for {
		result := <-chanIn
		if result == "begin" || result == "pong" {
			if player.Score >= 11 {
				fmt.Printf("%s Wins!!!Score %d\n", player.Name, player.Score)
				chanOut <- "pong"
				break
			}
			fmt.Printf("%s Hits \n", player.Name)
			chanOut <- "ping"
		}

		if result == "ping" {
			player.Score++
			fmt.Printf("%s hits \n", player.Name)
			if player.Score >= 11 {
				chanOut <- "stop"
				if player.Score >= 11 {
					fmt.Printf("%s Wins!!!Score %d\n", player.Name, player.Score)
					break
				}
			} else {
				chanOut <- "pong"
			}
		}
	}
}
