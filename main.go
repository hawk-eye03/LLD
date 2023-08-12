package main

import (
	"fmt"

	"github.com/hawk-eye03/LLD/TicTacToe/controllers"
	"github.com/hawk-eye03/LLD/TicTacToe/models"
)

func main() {
	// Initialize Game Controller
	gameController := controllers.GameController{}

	// Initialize number of players
	players := []models.IPlayer{
		models.NewPlayer("X", "P1", models.HUMAN),
		models.NewBot(models.EASY, models.NewPlayer("O", "P2", models.BOT)),
	}

	// Create a new Game with set of dimension, players & winning strategies
	game := gameController.CreateGame(3, players, nil)

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
			gameController.Undo()
		} else if input == "n" {
			gameController.Makemove(game)
		} else {
			fmt.Println("Invalid input!")
		}
	}

	gameController.PrintResult(*game)
}
