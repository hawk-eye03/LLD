package models

type Cell struct {
	row       int
	col       int
	cellState CellState
	player    Player
}

func (c *Cell) getRow() int {
	return c.row
}

func (c *Cell) getCol() int {
	return c.col
}

func (c *Cell) getCellState() CellState {
	return c.cellState
}

func (c *Cell) getPlayer() Player {
	return c.player
}

func (c *Cell) setPlayer(player Player) {
	c.player = player
}

func (c *Cell) setRow(row int) {
	c.row = row
}

func (c *Cell) setCol(col int) {
	c.col = col
}

func (c *Cell) setCellState(cellState CellState) {
	c.cellState = cellState
}
