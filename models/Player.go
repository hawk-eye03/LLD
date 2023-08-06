package models

type Player struct {
	symbol     string
	name       string
	id         int
	playerType PlayerType
}

func (p *Player) setSymbol(symbol string) {
	p.symbol = symbol
}

func (p *Player) setName(name string) {
	p.name = name
}

func (p *Player) setID(id int) {
	p.id = id
}

func (p *Player) setPlayerType(playerType PlayerType) {
	p.playerType = playerType
}

func (p *Player) getPlayerType() PlayerType {
	return p.playerType
}

func (p *Player) getID() int {
	return p.id
}

func (p *Player) getSymbol() string {
	return p.symbol
}

func (p *Player) getname() string {
	return p.name
}
