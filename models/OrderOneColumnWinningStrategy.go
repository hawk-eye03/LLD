package models

type OrderOneColumnWinningStrategy struct {
	columnMap []map[string]int
}

func getPlayerInfo(player IPlayer) *Player {
	var playerInfo Player
	switch p := player.(type) {
	case *Bot:
		playerInfo = p.GetBotPlayerInfo()
	case *Player:
		playerInfo = *p
	}
	return &playerInfo
}

func prepareMap(players []IPlayer) map[string]int {
	tempMap := map[string]int{}
	for j := 0; j < len(players); j++ {
		player := getPlayerInfo(players[j])
		tempMap[player.GetSymbol()] = 0
	}
	return tempMap
}
func prepareWinningStrategyMap(dimension int, players []IPlayer) []map[string]int {
	maps := []map[string]int{}
	for i := 0; i < dimension; i++ {
		tempMap := prepareMap(players)
		maps = append(maps, tempMap)
	}
	return maps
}

func NewOrderOneColumnWinningStrategy(dimension int, players []IPlayer) *OrderOneColumnWinningStrategy {
	return &OrderOneColumnWinningStrategy{
		columnMap: prepareWinningStrategyMap(dimension, players),
	}
}

func (o *OrderOneColumnWinningStrategy) CheckWinner(dimension int, move Move) bool {
	cell := move.GetCell()
	colMap := o.columnMap[cell.GetCol()]
	player := move.GetPlayer()
	colMap[player.GetSymbol()] += 1

	return colMap[player.GetSymbol()] == dimension
}

func (o *OrderOneColumnWinningStrategy) HandleUndo(dimension int, move Move) {
	cell := move.GetCell()
	colMap := o.columnMap[cell.GetCol()]
	player := move.GetPlayer()
	colMap[player.GetSymbol()] -= 1
}
