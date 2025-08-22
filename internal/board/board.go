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
	Turn Color
}

type Position struct {
	Row, Col int
}

type Move struct {
	StartPosition, EndPosition *Position
}

func NewBoard() *Board {
	board := &Board{
		Turn: Red,
	}
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

func (board *Board) GetColorString(color Color) (string, error) {
	switch color {
	case Empty:
		return "Empty", nil
	case Red:
		return "Red", nil
	case Black:
		return "Black", nil
	}
	return "", fmt.Errorf("No color associated with color")
}

func (board *Board) CheckValidMove(turn *Move) bool {
	return false
}

func (board *Board) ExecuteTurn(startRow, startCol, endRow, endCol int) error {
	startPosition := &Position{
		startRow,
		startCol,
	}
	endPosition := &Position{
		endRow,
		endCol,
	}
	turn := &Move{
		startPosition,
		endPosition,
	}
	if valid := board.CheckValidMove(turn); valid != true {
		return fmt.Errorf("Invalid move")
	}
	return nil
}
