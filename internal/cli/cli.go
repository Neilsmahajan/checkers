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
		fmt.Println("Input start and end positon in format A1,B2")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			input := scanner.Text()
			if input == "Q" || input == "q" {
				fmt.Println("Quitting the game...")
				return
			}
			if len(input) != 5 ||
				input[0] < 'A' || input[0] > 'H' ||
				input[1] < '1' || input[1] > '8' ||
				input[2] != ',' ||
				input[3] < 'A' || input[3] > 'H' ||
				input[4] < '1' || input[4] > '8' {
				fmt.Println("Invalid input")
				continue
			}
			fmt.Printf("%s", board.DrawBoard())
			fmt.Printf("Your input is %s\n", input)
			colorString, _ := board.GetColorString(board.Turn)
			fmt.Printf("It's %s's Turn\n", colorString)
			fmt.Println("Input start and end positon in format A1,B2")
		}
	}
}
