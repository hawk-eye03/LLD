package models

import "fmt"

type Board struct {
	size  int
	cells [][]Cell
}

func (b *Board) GetSize() int {
	return b.size
}

func (b *Board) SetSize(s int) {
	b.size = s
}

func (b *Board) GetCells() [][]Cell {
	return b.cells
}

func (b *Board) SetBoard(cells [][]Cell) {
	b.cells = cells
}

func NewBoard(dimension int) *Board {
	cells := [][]Cell{}
	for i := 0; i < dimension; i++ {
		tempCellRow := []Cell{}
		for j := 0; j < dimension; j++ {
			tempCellRow = append(tempCellRow, *NewCell(i, j))
		}
		cells = append(cells, tempCellRow)
	}
	return &Board{
		size:  dimension,
		cells: cells,
	}
}
func (b *Board) Print() {
	for i := range b.cells {
		fmt.Print("|")
		for j := range b.cells[i] {
			b.cells[i][j].Display()
		}
		fmt.Println()
	}
}
