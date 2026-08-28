package main

import (
	"github.com/khiyuazha-cell/Raid-3-Tic-Tac-Toe-Arena/internal/cli"
	"github.com/khiyuazha-cell/Raid-3-Tic-Tac-Toe-Arena/internal/game"
)

func main() {
	cfg := cli.Parse()
	game.Run(cfg)
}
