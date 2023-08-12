package models

import "fmt"

type Player struct {
	symbol     string
	name       string
	playerType PlayerType
}

func NewPlayer(symbol string, name string, playerType PlayerType) *Player {
	return &Player{
		symbol:     symbol,
		name:       name,
		playerType: playerType,
	}
}

func (p *Player) GetInfo() {
	fmt.Println("I am a Human")
}
func (p *Player) SetSymbol(symbol string) {
	p.symbol = symbol
}

func (p *Player) SetName(name string) {
	p.name = name
}

func (p *Player) SetPlayerType(playerType PlayerType) {
	p.playerType = playerType
}

func (p *Player) GetPlayerType() PlayerType {
	return p.playerType
}

func (p *Player) GetSymbol() string {
	return p.symbol
}

func (p *Player) GetName() string {
	return p.name
}
