package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/neilsmahajan/checkers/internal/board"
)

func Run(brd *board.Board) {
	for {
		fmt.Println("Checkers CLI")
		fmt.Printf("%s", brd.DrawBoard())
		colorString, _ := brd.GetColorString(brd.Turn)
		fmt.Printf("It's %s's Turn\n", colorString)
		fmt.Println("Input start and end position in format <Start Row><Start Column><End Row><End Column> (e.g. C1D2)")
		fmt.Println("To jump multiple pieces, chain positions (e.g. C1E3G5). Use Q to quit.")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			input := strings.TrimSpace(scanner.Text())
			if strings.EqualFold(input, "Q") {
				fmt.Println("Quitting the game...")
				return
			}
			if len(input) < 4 || len(input)%2 != 0 {
				fmt.Println("Invalid input length; must be even and at least 4 characters")
				continue
			}
			input = strings.ToUpper(input)

			// Mandatory capture enforcement (bug 5)
			if !strings.Contains(input, " ") { // simple heuristic; always check board
				if boardHasCapture := brd.PlayerHasCapture(brd.Turn); boardHasCapture {
					// Require that first move be a jump when capture exists
					// We'll detect if provided first segment is not a jump (row delta !=2)
					if len(input) == 4 { // single move provided
						startR, startC := int(input[0]-'A'), int(input[1]-'1')
						endR, endC := int(input[2]-'A'), int(input[3]-'1')
						if abs(endR-startR) != 2 || abs(endC-startC) != 2 {
							fmt.Println("You must take a capture if one is available")
							continue
						}
					}
				}
			}

			// Make a move slice
			moves := []*board.Move{}
			validInput := true
			for i := 0; i <= len(input)-4; i += 2 { // bug 2 fix: ensure i+3 in bounds
				if input[i] < 'A' || input[i] > 'H' ||
					input[i+1] < '1' || input[i+1] > '8' ||
					input[i+2] < 'A' || input[i+2] > 'H' ||
					input[i+3] < '1' || input[i+3] > '8' {
					fmt.Println("Invalid coordinate characters")
					validInput = false
					break
				}
				move := brd.NewMove(int(input[i]-'A'), int(input[i+1]-'1'), int(input[i+2]-'A'), int(input[i+3]-'1'))
				moves = append(moves, move)
			}
			if !validInput {
				continue
			}

			// Validate entire sequence first (bug 6)
			if err := brd.ValidateMoveSequence(moves); err != nil {
				fmt.Println("Invalid move sequence:", err)
				continue
			}

			for _, move := range moves {
				fmt.Printf("Move from %c%d to %c%d\n", move.StartPosition.Row+'A', move.StartPosition.Col+1, move.EndPosition.Row+'A', move.EndPosition.Col+1)
			}

			// Apply sequence after validation
			for _, move := range moves {
				if err := brd.ExecuteMove(move); err != nil {
					fmt.Println(err)
					validInput = false
					break
				}
			}
			if !validInput {
				continue
			}
			_ = brd.SwitchTurn()
			fmt.Printf("%s", brd.DrawBoard())
			colorString, _ := brd.GetColorString(brd.Turn)
			fmt.Printf("It's %s's Turn\n", colorString)
		}
	}
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
