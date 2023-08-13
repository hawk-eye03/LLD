package main

import (
	"fmt"

	"github.com/hawk-eye03/TicTacToe/controllers"
	"github.com/hawk-eye03/TicTacToe/models"
)

func main() {
	// Initialize Game Controller
	gameController := controllers.GameController{}

	// Initialize number of players
	players := []models.IPlayer{
		models.NewPlayer("X", "HUMANITY", models.HUMAN),
		models.NewBot(models.EASY, models.NewPlayer("O", "BOT", models.BOT)),
	}

	dimension := 3
	// Create a new Game with set of dimension, players & winning strategies
	game := gameController.CreateGame(dimension, players, []models.WinningStrategy{models.NewOrderOneColumnWinningStrategy(dimension, players), models.NewOrderOneRowWinningStrategy(dimension, players), models.NewOrderOneDiagonalWinningStrategy(dimension, players)})

	fmt.Println("\n------------------Game is starting------------------")

	for gameController.GetGameStatus(game) == models.IN_PROGRESS {
		// Print the current state of board
		fmt.Println("\nBelow is the current state of the game")
		gameController.DisplayBoard(game)

		// Ask user do you want to undo (y/n)?
		// if y -> undo
		// if n -> take input of row and col to make move
		var input string

		fmt.Println("Does anyone want to undo (y/n)")
		fmt.Scanln(&input)
		fmt.Println()

		if input == "y" {
			gameController.Undo(game)
		} else if input == "n" {
			gameController.Makemove(game)
		} else {
			fmt.Println("Invalid input!")
		}
	}

	gameController.PrintResult(*game)
	gameController.DisplayBoard(game)
}
