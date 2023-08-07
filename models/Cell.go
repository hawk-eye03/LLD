package models

import "fmt"

type Cell struct {
	row       int
	col       int
	cellState CellState
	player    Player
}

func (c *Cell) GetRow() int {
	return c.row
}

func (c *Cell) GetCol() int {
	return c.col
}

func (c *Cell) GetCellState() CellState {
	return c.cellState
}

func (c *Cell) GetPlayer() Player {
	return c.player
}

func (c *Cell) SetPlayer(player Player) {
	c.player = player
}

func (c *Cell) SetRow(row int) {
	c.row = row
}

func (c *Cell) SetCol(col int) {
	c.col = col
}

func (c *Cell) SetCellState(cellState CellState) {
	c.cellState = cellState
}

func (c *Cell) Display() {
	if c.cellState == EMPTY {
		fmt.Println(" - |")
	} else {
		fmt.Println(" " + c.player.symbol + " |")
	}
}
