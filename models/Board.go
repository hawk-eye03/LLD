package models

import "fmt"

type Board struct {
	size  int
	board [][]Cell
}

func (b *Board) GetSize() int {
	return b.size
}

func (b *Board) SetSize(s int) {
	b.size = s
}

func (b *Board) GetBoard() [][]Cell {
	return b.board
}

func (b *Board) SetBoard(board [][]Cell) {
	b.board = board
}

func (b *Board) Print() {
	// fmt.Println("Hello len of board =", len(b.board))
	for i := range b.board {
		fmt.Println("| ")
		for j := range b.board[i] {
			b.board[i][j].Display()
		}
	}
}
