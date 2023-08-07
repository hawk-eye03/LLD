package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/hawk-eye03/LLD/TicTacToe/controllers"
	"github.com/hawk-eye03/LLD/TicTacToe/models"
)

func main() {
	// Create a new Game
	gameController := controllers.GameController{}
	game := gameController.CreateGame()

	scanner := bufio.NewScanner(os.Stdin)

	for gameController.GetGameStatus() == models.IN_PROGRESS {
		// Print the board
		gameController.DisplayBoard(game)

		// Ask user do you want to undo (y/n)? (take input in terminal)
		// if y -> undo
		// if n -> take input of row and col to make move
		fmt.Println("Do you want to undo (y/n)")
		// Read the input from the user
		scanner.Scan()
		// Get the input as a string
		input := scanner.Text()

		if input == "y" {
			gameController.Undo()
		} else {
			// make move
			gameController.Makemove()
		}
	}

	// check status of game
	// if winner -> print winner
	// else print draw

	gameController.PrintResult(gameController)

}
