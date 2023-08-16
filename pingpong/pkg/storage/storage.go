package storage

import (
	"fmt"
	"sync"
	"thinknetica/pingpong/pkg/model"
)

type Player struct {
	storage map[string]*model.Player
	mu      sync.RWMutex
}

func New() *Player {
	return &Player{storage: make(map[string]*model.Player, 0)}
}

func (p *Player) Create(playerName string) *model.Player {
	defer p.mu.Unlock()
	p.mu.Lock()
	player := &model.Player{Name: playerName, Score: 0}
	p.storage[playerName] = player

	return player
}

func (p *Player) AddPoint(playerName string, point int) error {
	defer p.mu.Unlock()
	player, ok := p.storage[playerName]
	if !ok {
		return fmt.Errorf("player with name %s not exists", playerName)
	}

	player.Score = player.Score + point
	p.mu.Lock()
	p.storage[playerName] = player

	return nil
}

func (p *Player) GetPlayerByName(playerName string) (*model.Player, error) {
	player, ok := p.storage[playerName]
	if !ok {
		return nil, fmt.Errorf("player with name %s not exists", playerName)
	}

	return player, nil
}
