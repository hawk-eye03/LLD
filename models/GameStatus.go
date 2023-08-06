package models

type GameState int 

const(
	IN_PROGRESS = iota
	WIN
	DRAW
	PAUSE 
)