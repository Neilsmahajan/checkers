package board

import "fmt"

const (
	Width  = 8
	Height = 8
)

type Color int

const (
	Empty Color = iota
	Red
	Black
)

type Promotion int

const (
	Normal Promotion = iota
	Queen
)

type Piece struct {
	Color     Color
	Promotion Promotion
}

type Board struct {
	Grid [Height][Width]Piece
}

func NewBoard() *Board {
	board := &Board{}
	for row := range Height {
		for col := range Width {
			if row <= 2 && (row+col)%2 == 0 {
				board.Grid[row][col].Color = Red
			} else if row >= 5 && (row+col)%2 == 1 {
				board.Grid[row][col].Color = Black
			}
		}
	}
	return board
}

func (board *Board) DrawBoard() string {
	printedBoard := "    1    2    3    4    5    6    7    8"
	for row := range Height {
		printedBoard += fmt.Sprintf("\n  +----+----+----+----+----+----+----+----+\n%c |", 'A'+row)
		for col := range Width {
			switch board.Grid[row][col].Color {
			case Red:
				printedBoard += " 🔴 |"
			case Black:
				printedBoard += " ⚫️ |"
			case Empty:
				printedBoard += "    |"
			}
		}
	}
	printedBoard += "\n  +----+----+----+----+----+----+----+----+\n"
	return printedBoard
}
