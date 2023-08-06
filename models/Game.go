package models

import "github.com/hawk-eye03/LLD/TicTacToe/strategies/winningStrategies"

type Game struct {
	board             Board
	players           []Player
	moves             []Move
	winningStrategies []winningStrategies.WinningStrategy
	gameStatus        GameState
	nextMovePlayerInd int
	winner            Player
}

func (g *Game) setBoard(b Board) {
	g.board = b
}

func (g *Game) setPlayers(p []Player) {
	g.players = p
}

func (g *Game) setMoves(m []Move) {
	g.moves = m
}

func (g *Game) setWinningStrategies(winningStrategies []winningStrategies.WinningStrategy) {
	g.winningStrategies = winningStrategies
}

func (g *Game) setGameState(gameState GameState) {
	g.gameStatus = gameState
}

func (g *Game) setNextMovePlayerInd(ind int) {
	g.nextMovePlayerInd = ind
}

func (g *Game) setWinner(winner Player) {
	g.winner = winner
}

func (g *Game) getBoard() Board {
	return g.board
}

func (g *Game) getPlayers() []Player {
	return g.players
}

func (g *Game) getMoves() []Move {
	return g.moves
}

func (g *Game) getWinningStrategies() []winningStrategies.WinningStrategy {
	return g.winningStrategies
}

func (g *Game) getGameState() GameState {
	return g.gameStatus
}

func (g *Game) getNextMovePlayerInd() int {
	return g.nextMovePlayerInd
}

func (g *Game) getWinner() Player {
	return g.winner
}
