package models

type Move struct {
	player Player
	cell   Cell
}

func (m *Move) setPlayer(p Player) {
	m.player = p
}

func (m *Move) setCell(c Cell) {
	m.cell = c
}

func (m *Move) getPlayer() Player {
	return m.player
}

func (m *Move) getCell() Cell {
	return m.cell
}
