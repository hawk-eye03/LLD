package models

type Board struct{
	size int
	board [][]Cell
}

func (b *Board) getSize() int{
	return b.size
}

func (b *Board) setSize(s int){
	b.size = s
}

func (b *Board) getBoard() [][]Cell{
	return b.board
}

func (b *Board) setBoard(board [][]Cell){
	b.board = board
}



