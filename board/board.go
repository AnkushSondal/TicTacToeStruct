package board

import (
	"Assignment1/TicTacToeStruct/cell"
	"fmt"
)

type Board struct {
	Cells [9]*cell.Cell
}

func NewBoard() *Board {
	var cellBlocks [9]*cell.Cell
	for i := 0; i < 9; i++ {
		cellBlocks[i] = cell.NewCell()
	}
	return &Board{
		Cells: cellBlocks,
	}
}

func (b *Board) ResultAnalyzer() int {
	if b.Cells[0].ReturnMark() == b.Cells[1].ReturnMark() && b.Cells[1].ReturnMark() == b.Cells[2].ReturnMark() && b.Cells[0].IsCellMarked() { //row1
		return 1
	} else if b.Cells[3].ReturnMark() == b.Cells[4].ReturnMark() && b.Cells[4].ReturnMark() == b.Cells[5].ReturnMark() && b.Cells[3].IsCellMarked() { //row2
		return 1
	} else if b.Cells[6].ReturnMark() == b.Cells[7].ReturnMark() && b.Cells[7].ReturnMark() == b.Cells[8].ReturnMark() && b.Cells[6].IsCellMarked() { //row3
		return 1
	} else if b.Cells[0].ReturnMark() == b.Cells[3].ReturnMark() && b.Cells[3].ReturnMark() == b.Cells[6].ReturnMark() && b.Cells[0].IsCellMarked() { //column1
		return 1
	} else if b.Cells[1].ReturnMark() == b.Cells[4].ReturnMark() && b.Cells[4].ReturnMark() == b.Cells[7].ReturnMark() && b.Cells[1].IsCellMarked() { //column2
		return 1
	} else if b.Cells[2].ReturnMark() == b.Cells[5].ReturnMark() && b.Cells[5].ReturnMark() == b.Cells[8].ReturnMark() && b.Cells[2].IsCellMarked() { //column3
		return 1
	} else if b.Cells[0].ReturnMark() == b.Cells[4].ReturnMark() && b.Cells[4].ReturnMark() == b.Cells[8].ReturnMark() && b.Cells[0].IsCellMarked() { //diagonal1
		return 1
	} else if b.Cells[2].ReturnMark() == b.Cells[4].ReturnMark() && b.Cells[4].ReturnMark() == b.Cells[6].ReturnMark() && b.Cells[2].IsCellMarked() { //diagonal2
		return 1
	} else {
		return 0
	}

}

func (b *Board) ShowBoard() {
	for i := 0; i < 9; i++ {
		if i%3 == 0 {
			fmt.Println("\n-------------")
			fmt.Print("| ")
		}
		fmt.Print(b.Cells[i].ReturnMark(), " | ")
	}
	fmt.Print("\n-------------\n")
}
