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

func (p *Player) MakeMove(game *Game) Cell {
	var row, col int

	fmt.Print("Enter row (starting from 0): ")

	_, err1 := fmt.Scanln(&row)
	if err1 != nil {
		return *NewCell(-1, -1)
	}

	fmt.Print("Enter col (starting from 0): ")
	_, err2 := fmt.Scanln(&col)
	if err2 != nil {
		return *NewCell(-1, -1)
	}
	return *NewCell(row, col)
}
