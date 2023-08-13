package models

type BotPlayingStrategy interface {
	MakeMove(game *Game) Cell
}
