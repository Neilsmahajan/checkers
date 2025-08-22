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

func (board *Board) CheckValidMove(turn *Move) error {
	if board.Grid[turn.StartPosition.Row][turn.StartPosition.Col].Color != board.Turn {
		return fmt.Errorf("starting position doesn't correspond to a piece")
	}
	if board.Grid[turn.EndPosition.Row][turn.EndPosition.Col].Color != Empty {
		return fmt.Errorf("ending position must be empty")
	}
	return nil
}

func (board *Board) SwitchTurn() error {
	switch board.Turn {
	case Red:
		board.Turn = Black
	case Black:
		board.Turn = Red
	case Empty:
		return fmt.Errorf("turn isn't supposed to be blank")
	}
	return nil
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
	move := &Move{
		startPosition,
		endPosition,
	}
	if err := board.CheckValidMove(move); err != nil {
		return fmt.Errorf("invalid move: %v", err)
	}
	_ = board.SwitchTurn()
	return nil
}
