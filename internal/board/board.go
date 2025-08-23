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

func (board *Board) CheckIfPieceIsRightColorOrPromotionForDirection(move *Move, direction int) error {
	if board.Grid[move.StartPosition.Row][move.StartPosition.Col].Promotion == Queen {
		return nil
	}
	if (board.Turn == Red && direction == 1) || (board.Turn == Black && direction == -1) {
		return nil
	}
	return fmt.Errorf("your piece is not a queen and your move is in the wrong direction for the color")
}

func (board *Board) JumpOverOpponentPiece(opponentColor Color) error {
	switch board.Turn {
	case Red:
		if opponentColor != Black {
			return fmt.Errorf("a red piece must jump over a black piece")
		}
		return nil
	case Black:
		if opponentColor != Red {
			return fmt.Errorf("a black piece must jump over a red piece")
		}
	}
	return nil
}

func (board *Board) MovePiece(move *Move) error {
	if board.Grid[move.StartPosition.Row][move.StartPosition.Col].Color != board.Turn {
		return fmt.Errorf("starting position doesn't correspond to a piece of the right color")
	}
	if board.Grid[move.EndPosition.Row][move.EndPosition.Col].Color != Empty {
		return fmt.Errorf("ending position must be empty")
	}
	if (move.EndPosition.Row == move.StartPosition.Row-1) && (move.EndPosition.Col == move.StartPosition.Col-1 || move.EndPosition.Col == move.StartPosition.Col+1) {
		return board.CheckIfPieceIsRightColorOrPromotionForDirection(move, -1)
	}
	if (move.EndPosition.Row == move.StartPosition.Row+1) && (move.EndPosition.Col == move.StartPosition.Col-1 || move.EndPosition.Col == move.StartPosition.Col+1) {
		return board.CheckIfPieceIsRightColorOrPromotionForDirection(move, 1)
	}
	if (move.EndPosition.Row == move.StartPosition.Row-2) && (move.EndPosition.Col == move.StartPosition.Col-2) {
		if err := board.CheckIfPieceIsRightColorOrPromotionForDirection(move, -1); err != nil {
			return err
		}
		return board.JumpOverOpponentPiece(board.Grid[move.StartPosition.Row-1][move.EndPosition.Col-1].Color)
	}
	if (move.EndPosition.Row == move.StartPosition.Row-2) && (move.EndPosition.Col == move.StartPosition.Col+2) {
		if err := board.CheckIfPieceIsRightColorOrPromotionForDirection(move, -1); err != nil {
			return err
		}
		return board.JumpOverOpponentPiece(board.Grid[move.StartPosition.Row-1][move.EndPosition.Col+1].Color)

	}
	if (move.EndPosition.Row == move.StartPosition.Row+2) && (move.EndPosition.Col == move.StartPosition.Col-2) {
		if err := board.CheckIfPieceIsRightColorOrPromotionForDirection(move, 1); err != nil {
			return err
		}
		return board.JumpOverOpponentPiece(board.Grid[move.StartPosition.Row+1][move.EndPosition.Col-1].Color)
	}
	if (move.EndPosition.Row == move.StartPosition.Row+2) && (move.EndPosition.Col == move.StartPosition.Col+2) {
		if err := board.CheckIfPieceIsRightColorOrPromotionForDirection(move, 1); err != nil {
			return err
		}
		return board.JumpOverOpponentPiece(board.Grid[move.StartPosition.Row+1][move.EndPosition.Col+1].Color)
	}
	return fmt.Errorf("your move must move by one diagonal or two diagonals if taking an opponent piece")
}

func (board *Board) SwitchTurn() error {
	// ToDo: don't switch turn if piece takes opponent piece
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

	if err := board.MovePiece(move); err != nil {
		return fmt.Errorf("invalid move: %v", err)
	}
	_ = board.SwitchTurn()
	return nil
}
