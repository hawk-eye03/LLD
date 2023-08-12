package controllers

import (
	"github.com/hawk-eye03/LLD/TicTacToe/models"
	"github.com/hawk-eye03/LLD/TicTacToe/strategies/winningStrategies"
)

type GameController struct {
}

func (gc *GameController) CreateGame(dimension int, players []models.IPlayer, winningStrategies []winningStrategies.WinningStrategy) *models.Game {
	GameBuilder := &models.GameBuilder{}
	return GameBuilder.SetDimension(dimension).SetPlayers(players).SetWinningStrategies(winningStrategies).Build()
}

func (gc *GameController) DisplayBoard(game *models.Game) {
	game.PrintBoard()
}

func (gc *GameController) Undo() {

}

func (gc *GameController) Makemove(game *models.Game) {
	game.MakeMove()
}

func (gc *GameController) GetGameStatus(game *models.Game) models.GameState {
	return game.GetGameState()
}

// check status of game
// if winner -> print winner
// else print draw
func (gc *GameController) PrintResult(game models.Game) {
	game.PrintResult()
}
