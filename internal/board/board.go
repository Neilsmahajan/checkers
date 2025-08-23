package board

import (
	"fmt"
)

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
			} else if row >= 5 && (row+col)%2 == 0 {
				board.Grid[row][col].Color = Black
			}
		}
	}
	return board
}

func (board *Board) NewMove(startRow, startCol, endRow, endCol int) *Move {
	return &Move{
		StartPosition: &Position{
			Row: startRow,
			Col: startCol,
		},
		EndPosition: &Position{
			Row: endRow,
			Col: endCol,
		},
	}
}

func (board *Board) IsValidJumpMove(move *Move) error {
	var err error
	if (move.EndPosition.Row == move.StartPosition.Row-2) && (move.EndPosition.Col == move.StartPosition.Col-2) {
		err = board.CheckIfPieceIsRightColorOrPromotionForDirection(move, -1)
	} else if (move.EndPosition.Row == move.StartPosition.Row-2) && (move.EndPosition.Col == move.StartPosition.Col+2) {
		err = board.CheckIfPieceIsRightColorOrPromotionForDirection(move, -1)
	} else if (move.EndPosition.Row == move.StartPosition.Row+2) && (move.EndPosition.Col == move.StartPosition.Col-2) {
		err = board.CheckIfPieceIsRightColorOrPromotionForDirection(move, 1)
	} else if (move.EndPosition.Row == move.StartPosition.Row+2) && (move.EndPosition.Col == move.StartPosition.Col+2) {
		err = board.CheckIfPieceIsRightColorOrPromotionForDirection(move, 1)
	} else {
		return fmt.Errorf("your jump move must move by two diagonals for multiple jumps")
	}
	return err
}

func (board *Board) DrawBoard() string {
	printedBoard := "    1    2    3    4    5    6    7    8"
	for row := range Height {
		printedBoard += fmt.Sprintf("\n  +----+----+----+----+----+----+----+----+\n%c |", 'A'+row)
		for col := range Width {
			switch board.Grid[row][col] {
			case Piece{Color: Red, Promotion: Normal}:
				printedBoard += " 🔴 |"
			case Piece{Color: Red, Promotion: Queen}:
				printedBoard += " 🟥 |"
			case Piece{Color: Black, Promotion: Normal}:
				printedBoard += " ⚫️ |"
			case Piece{Color: Black, Promotion: Queen}:
				printedBoard += " ⬛️ |"
			case Piece{Color: Empty, Promotion: Normal}:
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

func (board *Board) MakeQueenIfPossible(move *Move) {
	if board.Turn == Red && move.EndPosition.Row == 7 {
		board.Grid[move.EndPosition.Row][move.EndPosition.Col].Promotion = Queen
	}
	if board.Turn == Black && move.EndPosition.Row == 0 {
		board.Grid[move.EndPosition.Row][move.EndPosition.Col].Promotion = Queen
	}
}

func (board *Board) MoveOneDiagonal(move *Move) error {
	tempPiece := board.Grid[move.StartPosition.Row][move.StartPosition.Col]
	board.Grid[move.StartPosition.Row][move.StartPosition.Col].Color = Empty
	board.Grid[move.StartPosition.Row][move.StartPosition.Col].Promotion = Normal
	board.Grid[move.EndPosition.Row][move.EndPosition.Col] = tempPiece
	board.MakeQueenIfPossible(move)
	return nil
}

func (board *Board) JumpOverOpponentPiece(move *Move, jumpedPiece *Piece) error {
	if board.Turn == Red && jumpedPiece.Color != Black {
		return fmt.Errorf("a red piece must jump over a black piece")
	}
	if board.Turn == Black && jumpedPiece.Color != Red {
		return fmt.Errorf("a black piece must jump over a red piece")
	}
	tempPiece := board.Grid[move.StartPosition.Row][move.StartPosition.Col]
	board.Grid[move.StartPosition.Row][move.StartPosition.Col].Color = Empty
	board.Grid[move.StartPosition.Row][move.StartPosition.Col].Promotion = Normal
	board.Grid[move.EndPosition.Row][move.EndPosition.Col] = tempPiece
	jumpedPiece.Color = Empty
	jumpedPiece.Promotion = Normal
	board.MakeQueenIfPossible(move)
	return nil
}

func (board *Board) MovePiece(move *Move) error {
	// Bounds safety (bug 3)
	if move.StartPosition.Row < 0 || move.StartPosition.Row >= Height ||
		move.StartPosition.Col < 0 || move.StartPosition.Col >= Width ||
		move.EndPosition.Row < 0 || move.EndPosition.Row >= Height ||
		move.EndPosition.Col < 0 || move.EndPosition.Col >= Width {
		return fmt.Errorf("move out of bounds")
	}
	if board.Grid[move.StartPosition.Row][move.StartPosition.Col].Color != board.Turn {
		return fmt.Errorf("starting position doesn't correspond to a piece of the right color")
	}
	if board.Grid[move.EndPosition.Row][move.EndPosition.Col].Color != Empty {
		return fmt.Errorf("ending position must be empty")
	}

	// Move by one diagonal
	if (move.EndPosition.Row == move.StartPosition.Row-1) && (move.EndPosition.Col == move.StartPosition.Col-1 || move.EndPosition.Col == move.StartPosition.Col+1) {
		if err := board.CheckIfPieceIsRightColorOrPromotionForDirection(move, -1); err != nil {
			return err
		}
		return board.MoveOneDiagonal(move)
	}
	if (move.EndPosition.Row == move.StartPosition.Row+1) && (move.EndPosition.Col == move.StartPosition.Col-1 || move.EndPosition.Col == move.StartPosition.Col+1) {
		if err := board.CheckIfPieceIsRightColorOrPromotionForDirection(move, 1); err != nil {
			return err
		}
		return board.MoveOneDiagonal(move)
	}

	// Move by two diagonals (jump over opponent piece)
	if (move.EndPosition.Row == move.StartPosition.Row-2) && (move.EndPosition.Col == move.StartPosition.Col-2) {
		if err := board.CheckIfPieceIsRightColorOrPromotionForDirection(move, -1); err != nil {
			return err
		}
		return board.JumpOverOpponentPiece(move, &board.Grid[move.StartPosition.Row-1][move.StartPosition.Col-1])
	}
	if (move.EndPosition.Row == move.StartPosition.Row-2) && (move.EndPosition.Col == move.StartPosition.Col+2) {
		if err := board.CheckIfPieceIsRightColorOrPromotionForDirection(move, -1); err != nil {
			return err
		}
		return board.JumpOverOpponentPiece(move, &board.Grid[move.StartPosition.Row-1][move.StartPosition.Col+1])

	}
	if (move.EndPosition.Row == move.StartPosition.Row+2) && (move.EndPosition.Col == move.StartPosition.Col-2) {
		if err := board.CheckIfPieceIsRightColorOrPromotionForDirection(move, 1); err != nil {
			return err
		}
		return board.JumpOverOpponentPiece(move, &board.Grid[move.StartPosition.Row+1][move.StartPosition.Col-1])
	}
	if (move.EndPosition.Row == move.StartPosition.Row+2) && (move.EndPosition.Col == move.StartPosition.Col+2) {
		if err := board.CheckIfPieceIsRightColorOrPromotionForDirection(move, 1); err != nil {
			return err
		}
		return board.JumpOverOpponentPiece(move, &board.Grid[move.StartPosition.Row+1][move.StartPosition.Col+1])
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

func (board *Board) ExecuteMove(move *Move) error {
	if err := board.MovePiece(move); err != nil {
		return fmt.Errorf("invalid move: %v", err)
	}
	return nil
}

// PlayerHasCapture returns true if the current player (color) has at least one legal capture available (bug 5)
func (board *Board) PlayerHasCapture(color Color) bool {
	opponent := func(c Color) Color {
		if c == Red {
			return Black
		}
		if c == Black {
			return Red
		}
		return Empty
	}
	opp := opponent(color)
	for r := 0; r < Height; r++ {
		for c := 0; c < Width; c++ {
			piece := board.Grid[r][c]
			if piece.Color != color {
				continue
			}
			// Directions a non-queen can move
			dirs := [][2]int{{2, 2}, {2, -2}, {-2, 2}, {-2, -2}}
			for _, d := range dirs {
				dr, dc := d[0], d[1]
				// Restrict non-queens to forward direction
				if piece.Promotion != Queen {
					if color == Red && dr < 0 { // Red moves increasing row
						continue
					}
					if color == Black && dr > 0 { // Black moves decreasing row
						continue
					}
				}
				r2, c2 := r+dr, c+dc
				rMid, cMid := r+dr/2, c+dc/2
				if r2 < 0 || r2 >= Height || c2 < 0 || c2 >= Width || rMid < 0 || rMid >= Height || cMid < 0 || cMid >= Width {
					continue
				}
				if board.Grid[r2][c2].Color == Empty && board.Grid[rMid][cMid].Color == opp {
					return true
				}
			}
		}
	}
	return false
}

// ValidateMoveSequence validates a sequence of moves before applying (bug 6)
// Ensures continuity, bounds, jump correctness (when len>1), and that each move is legal.
func (board *Board) ValidateMoveSequence(moves []*Move) error {
	if len(moves) == 0 {
		return fmt.Errorf("no moves provided")
	}
	// Copy board for dry run
	temp := *board
	// For multi-move sequence, all must be jumps and contiguous
	if len(moves) > 1 {
		for i, m := range moves {
			if i > 0 { // continuity: start equals previous end
				prev := moves[i-1]
				if m.StartPosition.Row != prev.EndPosition.Row || m.StartPosition.Col != prev.EndPosition.Col {
					return fmt.Errorf("move %d does not start where previous move ended", i+1)
				}
			}
			dr := m.EndPosition.Row - m.StartPosition.Row
			dc := m.EndPosition.Col - m.StartPosition.Col
			if (dr != 2 && dr != -2) || (dc != 2 && dc != -2) {
				return fmt.Errorf("all moves in a multi-move sequence must be jumps")
			}
			if err := temp.MovePiece(m); err != nil { // will validate capture legality
				return err
			}
		}
		return nil
	}
	// Single move: just validate by simulation
	if err := temp.MovePiece(moves[0]); err != nil {
		return err
	}
	return nil
}

func (board *Board) CheckWinCondition() (Color, bool) {
	redCount, blackCount := 0, 0
	for r := range Height {
		for c := range Width {
			switch board.Grid[r][c].Color {
			case Red:
				redCount++
			case Black:
				blackCount++
			}
		}
	}
	if redCount == 0 {
		return Black, true
	}
	if blackCount == 0 {
		return Red, true
	}
	return Empty, false
}
