package storage

import (
	"fmt"
	"sync"
)

type Player struct {
	storage map[string]int
	mu      sync.RWMutex
}

func New() *Player {
	return &Player{storage: make(map[string]int, 0)}
}

func (p *Player) Create(playerName string) {
	defer p.mu.Unlock()
	p.mu.Lock()
	p.storage[playerName] = 0
}

func (p *Player) AddPoint(playerName string, point int) error {
	defer p.mu.Unlock()
	score, ok := p.storage[playerName]
	if !ok {
		return fmt.Errorf("player with name %s not exists", playerName)
	}

	score = score + point
	p.mu.Lock()
	p.storage[playerName] = score

	return nil
}

func (p *Player) GetPlayerScore(playerName string) (int, error) {
	score, ok := p.storage[playerName]
	if !ok {
		return 0, fmt.Errorf("player with name %s not exists", playerName)
	}

	return score, nil
}
