package models

type WinningStrategy interface {
	CheckWinner(dimension int, move Move) bool
	HandleUndo(dimension int, move Move)
}
