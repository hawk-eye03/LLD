package models

import (
	"fmt"
)

type Game struct {
	board                Board
	players              []IPlayer
	moves                []Move
	winningStrategies    []WinningStrategy
	gameState            GameState
	currentMovePlayerInd int
	winner               Player
}

func newGame(builder *GameBuilder) *Game {
	return &Game{
		board:                *NewBoard(builder.dimension),
		players:              builder.players,
		winningStrategies:    builder.winningStrategies,
		gameState:            IN_PROGRESS,
		currentMovePlayerInd: 0,
	}
}

type GameBuilder struct {
	dimension         int
	players           []IPlayer
	winningStrategies []WinningStrategy
}

func (g *GameBuilder) SetDimension(dimension int) *GameBuilder {
	g.dimension = dimension
	return g
}

func (g *GameBuilder) SetPlayers(players []IPlayer) *GameBuilder {
	g.players = players
	return g
}

func (g *GameBuilder) SetWinningStrategies(winningStrategies []WinningStrategy) *GameBuilder {
	g.winningStrategies = winningStrategies
	return g
}

func appendAndCheckHmSymbol(hm *map[string]bool, p Player) bool {
	if (*hm)[p.symbol] {
		return false
	}
	(*hm)[p.symbol] = true
	return true
}
func (g *GameBuilder) validate() bool {
	// check number of players should be less than dimension of board
	noOfPlayers := len(g.players)
	if noOfPlayers > g.dimension-1 {
		return false
	}
	// check number of players cannot be less than 2
	if noOfPlayers < 2 {
		return false
	}
	botCount := 0
	for _, player := range g.players {
		switch player.(type) {
		case *Bot:
			botCount += 1
		}
	}
	// check there is at max one bot playing
	if botCount > 1 {
		return false
	}

	// check two players dont have same symbol
	hmSymbol := map[string]bool{}
	for _, player := range g.players {
		switch p := player.(type) {
		case *Bot:
			if !appendAndCheckHmSymbol(&hmSymbol, p.player) {
				return false
			}
		case *Player:
			if !appendAndCheckHmSymbol(&hmSymbol, *p) {
				return false
			}
		}
	}
	return true
}
func (g *GameBuilder) Build() *Game {
	if !g.validate() {
		fmt.Println("Invalid parameters for creating game")
	}
	return newGame(g)
}

func (g *Game) SetBoard(b Board) {
	g.board = b
}

func (g *Game) SetMoves(m []Move) {
	g.moves = m
}

func (g *Game) SetWinningStrategies(winningStrategies []WinningStrategy) {
	g.winningStrategies = winningStrategies
}

func (g *Game) SetGameState(gameState GameState) {
	g.gameState = gameState
}

func (g *Game) SetCurrMovePlayerInd(ind int) {
	g.currentMovePlayerInd = ind
}

func (g *Game) SetWinner(winner Player) {
	g.winner = winner
}

func (g *Game) GetBoard() Board {
	return g.board
}

func (g *Game) GetPlayers() []IPlayer {
	return g.players
}

func (g *Game) GetMoves() []Move {
	return g.moves
}

func (g *Game) GetWinningStrategies() []WinningStrategy {
	return g.winningStrategies
}

func (g *Game) GetGameState() GameState {
	return g.gameState
}

func (g *Game) GetCurrMovePlayerInd() int {
	return g.currentMovePlayerInd
}

func (g *Game) GetWinner() Player {
	return g.winner
}

func (g *Game) PrintBoard() {
	g.board.Print()
}

func (g *Game) PrintResult() {
	gameState := g.GetGameState()
	if gameState == ENDED {
		fmt.Println("\nGame has Ended!")
		fmt.Println("Winner is: ", g.winner.GetName())
		fmt.Println()
	} else {
		fmt.Println("\nGame Drawn!")
		fmt.Println()
	}
}

// Display player name who has the current turn
func (g *Game) GetPlayerAndDisplayTurn() Player {
	var playerInfo *Player

	players := g.GetPlayers()
	currentPlayer := players[g.GetCurrMovePlayerInd()]
	playerInfo = getPlayerInfo(currentPlayer)

	fmt.Println("It is " + playerInfo.GetName() + "'s turn")

	return *playerInfo
}

func (g *Game) MakeMove() {
	currPlayer := g.GetPlayerAndDisplayTurn()
	currentPlayer := g.players[g.GetCurrMovePlayerInd()]
	proposedCell := currentPlayer.MakeMove(g)

	row := proposedCell.GetRow()
	col := proposedCell.GetCol()

	// check if the move given by player (Human/Bot) is valid or not
	if !(g.ValidateMove(row, col)) || row == -1 && col == -1 {
		fmt.Println("Invalid move. Retry")
		return
	}
	fmt.Println("\nMove was made at row:", row, ", col:", col)

	board := g.GetBoard()
	cells := board.GetCells()

	cells[row][col].SetCellState(FILLED)
	cells[row][col].SetPlayer(currPlayer)

	move := NewMove(currPlayer, &cells[row][col])

	// append moves to list of moves in game
	g.moves = append(g.moves, *move)

	// check whether game is won or drawn after the move is complete
	if g.CheckGameWon(*move) {
		return
	}
	if g.CheckGameDrawn() {
		return
	}

	// set next player index
	g.SetCurrMovePlayerInd((g.GetCurrMovePlayerInd() + 1) % len(g.players))
}

func (g *Game) UndoMove() {
	if len(g.moves) == 0 {
		fmt.Println("No moves yet. Can't Undo.")
		return
	}
	// remove last move
	lastMove := g.moves[len(g.moves)-1]
	lastCell := lastMove.GetCell()
	g.moves = g.moves[:len(g.moves)-1]
	g.currentMovePlayerInd = (g.currentMovePlayerInd - 1 + len(g.players)) % len(g.players)
	lastCell.SetCellState(EMPTY)

	for _, winningStrategy := range g.winningStrategies {
		winningStrategy.HandleUndo(g.board.GetSize(), lastMove)
	}
}

func (g *Game) CheckGameWon(move Move) bool {
	for _, winningStrategy := range g.GetWinningStrategies() {
		board := g.GetBoard()
		if winningStrategy.CheckWinner(board.GetSize(), move) {
			g.gameState = ENDED
			g.winner = move.GetPlayer()
			return true
		}
	}
	return false
}

func (g *Game) CheckGameDrawn() bool {
	if len(g.moves) == (g.GetBoard().size)*(g.GetBoard().size) {
		g.gameState = DRAW
		return true
	}
	return false
}

func (g *Game) ValidateMove(row, col int) bool {
	if row >= g.board.size || row < 0 {
		return false
	}
	if col >= g.board.size || col < 0 {
		return false
	}

	board := g.GetBoard()
	cells := board.GetCells()

	return cells[row][col].cellState == EMPTY
}
