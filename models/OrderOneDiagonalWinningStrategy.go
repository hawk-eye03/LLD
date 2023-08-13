package models

type OrderOneDiagonalWinningStrategy struct {
	leftDiagMap  map[string]int
	rightDiagMap map[string]int
}

func NewOrderOneDiagonalWinningStrategy(dimension int, players []IPlayer) *OrderOneDiagonalWinningStrategy {
	return &OrderOneDiagonalWinningStrategy{
		leftDiagMap:  prepareMap(players),
		rightDiagMap: prepareMap(players),
	}
}

func (o *OrderOneDiagonalWinningStrategy) CheckWinner(dimension int, move Move) bool {

	cell := move.GetCell()
	currPlayer := move.GetPlayer()
	row := cell.GetRow()
	col := cell.GetCol()
	if (row != col) && (row+col) != dimension-1 {
		return false
	}
	if row == col {
		o.leftDiagMap[currPlayer.GetSymbol()] += 1
	}

	if row+col == dimension-1 {
		o.rightDiagMap[currPlayer.GetSymbol()] += 1
	}

	if o.rightDiagMap[currPlayer.GetSymbol()] == dimension || o.leftDiagMap[currPlayer.GetSymbol()] == dimension {
		return true
	}
	return false
}

func (o *OrderOneDiagonalWinningStrategy) HandleUndo(dimension int, move Move) {
	cell := move.GetCell()
	currPlayer := move.GetPlayer()
	row := cell.GetRow()
	col := cell.GetCol()
	if row == col {
		o.leftDiagMap[currPlayer.GetSymbol()] -= 1
	}
	if row+col == dimension-1 {
		o.rightDiagMap[currPlayer.GetSymbol()] -= 1
	}
}
