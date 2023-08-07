package controllers

import (
	"fmt"

	"github.com/hawk-eye03/LLD/TicTacToe/models"
)

type GameController struct {
}

func (gc *GameController) CreateGame() models.Game {
	return models.Game{}
}

func (gc *GameController) PrintBoard() {

}

func (gc *GameController) DisplayBoard(game models.Game) {
	game.PrintBoard()
}

func (gc *GameController) Undo() {

}

func (gc *GameController) Makemove() {

}

func (gc *GameController) GetGameStatus() int {
	return 0
}

func (gc *GameController) PrintWinner() {

}

func (gc *GameController) PrintResult(gameController GameController) {
	gameStatus := gameController.GetGameStatus()
	if gameStatus == models.WIN {
		gameController.PrintWinner()
	} else {
		fmt.Println("Game Drawn!")
	}
}
