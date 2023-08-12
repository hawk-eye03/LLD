package models

type Move struct {
	player Player
	cell   Cell
}

func NewMove(player Player, cell Cell) *Move {
	return &Move{
		player: player,
		cell:   cell,
	}
}

func (m *Move) SetPlayer(p Player) {
	m.player = p
}

func (m *Move) SetCell(c Cell) {
	m.cell = c
}

func (m *Move) GetPlayer() Player {
	return m.player
}

func (m *Move) GetCell() Cell {
	return m.cell
}
