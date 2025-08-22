package cli

import (
	"fmt"

	"github.com/neilsmahajan/checkers/internal/board"
)

func Run() {
	board := board.NewBoard()
	fmt.Printf("%s", board.DrawBoard())
}
