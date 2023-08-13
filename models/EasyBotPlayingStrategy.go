package models

type EasyBotPlayingStrategy struct{}

func (e *EasyBotPlayingStrategy) MakeMove(game *Game) Cell {
	board := game.GetBoard()
	cells := board.GetCells()
	for i := range cells {
		for j := 0; j < board.GetSize(); j++ {
			if cells[i][j].GetCellState() == EMPTY {
				return cells[i][j]
			}
		}
	}
	return *NewCell(-1, -1)
}
