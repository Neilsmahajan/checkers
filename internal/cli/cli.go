package cli

import (
	"bufio"
	"fmt"
	"os"

	"github.com/neilsmahajan/checkers/internal/board"
)

func Run(brd *board.Board) {
	for {
		fmt.Println("Checkers CLI")
		fmt.Printf("%s", brd.DrawBoard())
		colorString, _ := brd.GetColorString(brd.Turn)
		fmt.Printf("It's %s's Turn\n", colorString)
		fmt.Println("Input start and end positon in format <Start Row><Start Column><End Row><End Column>, (e.g. C1D2)")
		fmt.Println("If you want to jump multiple pieces, input the positions one after another (e.g. C1E3G5)")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			input := scanner.Text()
			if input == "Q" || input == "q" {
				fmt.Println("Quitting the game...")
				return
			}

			// Make a move slice
			moves := []*board.Move{}
			for i := 0; i < len(input)-2; i += 2 {
				if input[i] < 'A' || input[i] > 'H' ||
					input[i+1] < '1' || input[i+1] > '8' ||
					input[i+2] < 'A' || input[i+2] > 'H' ||
					input[i+3] < '1' || input[i+3] > '8' {
					fmt.Println("Invalid input")
					continue
				}
				move := brd.NewMove(int(input[i]-'A'), int(input[i+1]-'1'), int(input[i+2]-'A'), int(input[i+3]-'1'))
				moves = append(moves, move)
			}

			for _, move := range moves {
				fmt.Printf("Move from %c%d to %c%d\n", move.StartPosition.Row+'A', move.StartPosition.Col+1, move.EndPosition.Row+'A', move.EndPosition.Col+1)
			}

			validMoves := true
			if len(moves) > 1 {
				for _, move := range moves {
					if err := brd.IsValidJumpMove(move); err != nil {
						fmt.Println(err)
						validMoves = false
						break
					}
				}
			}

			if validMoves {
				for _, move := range moves {
					if err := brd.ExecuteMove(move); err != nil {
						fmt.Println(err)
						validMoves = false
						break
					}
				}
			}
			if validMoves {
				_ = brd.SwitchTurn()
			}
			fmt.Printf("%s", brd.DrawBoard())
			colorString, _ := brd.GetColorString(brd.Turn)
			fmt.Printf("It's %s's Turn\n", colorString)
		}
	}
}
