package models

type OrderOneRowWinningStrategy struct {
	rowMap []map[string]int
}

func NewOrderOneRowWinningStrategy(dimension int, players []IPlayer) *OrderOneRowWinningStrategy {
	return &OrderOneRowWinningStrategy{
		rowMap: prepareWinningStrategyMap(dimension, players),
	}
}

func (o *OrderOneRowWinningStrategy) CheckWinner(dimension int, move Move) bool {
	cell := move.GetCell()
	rowMap := o.rowMap[cell.GetRow()]
	player := move.GetPlayer()
	rowMap[player.GetSymbol()] += 1

	return rowMap[player.GetSymbol()] == dimension
}

func (o *OrderOneRowWinningStrategy) HandleUndo(dimension int, move Move) {
	cell := move.GetCell()
	rowMap := o.rowMap[cell.GetRow()]
	player := move.GetPlayer()
	rowMap[player.GetSymbol()] -= 1
}
