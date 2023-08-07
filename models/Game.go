package models

import (
	"github.com/hawk-eye03/LLD/TicTacToe/strategies/winningStrategies"
)

type Game struct {
	board             Board
	players           []Player
	moves             []Move
	winningStrategies []winningStrategies.WinningStrategy
	gameStatus        GameState
	nextMovePlayerInd int
	winner            Player
}

func (g *Game) SetBoard(b Board) {
	g.board = b
}

func (g *Game) SetPlayers(p []Player) {
	g.players = p
}

func (g *Game) SetMoves(m []Move) {
	g.moves = m
}

func (g *Game) SetWinningStrategies(winningStrategies []winningStrategies.WinningStrategy) {
	g.winningStrategies = winningStrategies
}

func (g *Game) SetGameState(gameState GameState) {
	g.gameStatus = gameState
}

func (g *Game) SetNextMovePlayerInd(ind int) {
	g.nextMovePlayerInd = ind
}

func (g *Game) SetWinner(winner Player) {
	g.winner = winner
}

func (g *Game) GetBoard() Board {
	return g.board
}

func (g *Game) GetPlayers() []Player {
	return g.players
}

func (g *Game) GetMoves() []Move {
	return g.moves
}

func (g *Game) GetWinningStrategies() []winningStrategies.WinningStrategy {
	return g.winningStrategies
}

func (g *Game) GetGameState() GameState {
	return g.gameStatus
}

func (g *Game) GetNextMovePlayerInd() int {
	return g.nextMovePlayerInd
}

func (g *Game) GetWinner() Player {
	return g.winner
}

func (g *Game) PrintBoard() {
	g.board.Print()
}
