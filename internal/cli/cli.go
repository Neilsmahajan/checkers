package cli

import (
	"bufio"
	"fmt"
	"os"

	"github.com/neilsmahajan/checkers/internal/board"
)

func Run(board *board.Board) {
	for {
		fmt.Println("Checkers CLI")
		fmt.Printf("%s", board.DrawBoard())
		colorString, _ := board.GetColorString(board.Turn)
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
			// if input[0] < 'A' || input[0] > 'H' ||
			// 	input[1] < '1' || input[1] > '8' ||
			// 	input[2] < 'A' || input[2] > 'H' ||
			// 	input[3] < '1' || input[3] > '8' {
			// 	fmt.Println("Invalid input")
			// 	continue
			// }

			// if err := board.ExecuteTurn(int(input[0]-'A'), int(input[1]-'1'), int(input[2]-'A'), int(input[3]-'1')); err != nil {
			// 	fmt.Printf("Error executing turn %v\n", err)
			// }

			for i := 0; i < len(input)-2; i += 2 {
				if input[i] < 'A' || input[i] > 'H' ||
					input[i+1] < '1' || input[i+1] > '8' ||
					input[i+2] < 'A' || input[i+2] > 'H' ||
					input[i+3] < '1' || input[i+3] > '8' {
					fmt.Println("Invalid input")
					continue
				}
				if err := board.ExecuteTurn(int(input[i]-'A'), int(input[i+1]-'1'), int(input[i+2]-'A'), int(input[i+3]-'1')); err != nil {
					fmt.Printf("Error executing turn %v\n", err)
				}
			}
			_ = board.SwitchTurn()
			fmt.Printf("%s", board.DrawBoard())
			colorString, _ := board.GetColorString(board.Turn)
			fmt.Printf("It's %s's Turn\n", colorString)
		}
	}
}
