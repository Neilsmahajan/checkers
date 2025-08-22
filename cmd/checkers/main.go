package main

import (
	"github.com/neilsmahajan/checkers/internal/board"
	"github.com/neilsmahajan/checkers/internal/cli"
)

func main() {
	cli.Run(board.NewBoard())
}
