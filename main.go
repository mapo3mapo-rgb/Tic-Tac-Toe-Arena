package main

import (
	"github.com/mapo3mapo-rgb/Tic-Tac-Toe-Arena/internal/cli"
	"github.com/mapo3mapo-rgb/Tic-Tac-Toe-Arena/internal/game"
)

func main() {
	cfg := cli.Parse()
	game.Run(cfg)
}
