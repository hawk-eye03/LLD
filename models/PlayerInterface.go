package models

type IPlayer interface {
	MakeMove(game *Game) Cell
}
