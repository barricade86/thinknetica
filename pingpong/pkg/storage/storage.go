package storage

import (
	"fmt"
	"sync"
	"thinknetica/pingpong/pkg/model"
)

type Match struct {
	storage map[string]*model.Player
	mu      sync.RWMutex
}

func New() *Match {
	return &Match{storage: make(map[string]*model.Player, 0)}
}

func (ms *Match) CreatePlayer(playerName string) {
	defer ms.mu.Unlock()
	ms.mu.Lock()
	player := &model.Player{Name: playerName, Score: 0}
	ms.storage[playerName] = player
}

func (ms *Match) AddPoint(playerName string, point uint) error {
	defer ms.mu.Unlock()
	player, ok := ms.storage[playerName]
	if !ok {
		return fmt.Errorf("player with name %s not exists", playerName)
	}

	player.Score = player.Score + point
	ms.mu.Lock()
	ms.storage[playerName] = player

	return nil
}

func (ms *Match) GetPlayerByName(playerName string) (*model.Player, error) {
	player, ok := ms.storage[playerName]
	if !ok {
		return nil, fmt.Errorf("player with name %s not exists", playerName)
	}

	return player, nil
}

func (ms *Match) GetTotalScore() map[string]uint {
	score := make(map[string]uint, 0)
	for key, val := range ms.storage {
		score[key] = val.Score
	}

	return score
}
