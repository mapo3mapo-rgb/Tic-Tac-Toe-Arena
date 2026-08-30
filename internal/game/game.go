package game

import (
	"fmt"
	"os"

	"github.com/mapo3mapo-rgb/Tic-Tac-Toe-Arena/internal/ai"
	"github.com/mapo3mapo-rgb/Tic-Tac-Toe-Arena/internal/board"
	"github.com/mapo3mapo-rgb/Tic-Tac-Toe-Arena/internal/cli"
)

type Stats struct {
	Games     int
	WinsX     int
	WinsO     int
	Draws     int
	MovesThis int
}

func Run(cfg cli.Config) {
	b := board.New(cfg.Size)
	stats := Stats{}

	for {
		b.Reset()
		stats.MovesThis = 0
		current := cfg.First
		playGame(&b, cfg, current, &stats)

		printStats(stats, cfg.Verbose)

		fmt.Print("Play again? (y/n): ")
		var answer string
		fmt.Scan(&answer)
		if answer != "y" {
			break
		}
	}
}

func playGame(b *board.Board, cfg cli.Config, first string, stats *Stats) {
	current := first

	for {
		b.PrintWithOptions(cfg.Color, cfg.Big, nil) // ← было b.Print()

		var cell int

		if cfg.Mode == "ai" && current == "O" {
			var reason string
			cell, reason = ai.GetMove(*b)

			fmt.Printf("O plays %d\n", cell)

			if cfg.Verbose {
				fmt.Printf("AI: %s at %d\n", reason, cell)
			}
		} else {
			name := cfg.NameX
			if current == "O" {
				name = cfg.NameO
			}
			cell = getInput(b, name, cfg.Size)
		}

		b.Place(cell, current)
		stats.MovesThis++

		win, winLine := b.CheckWin(current) // ← было win, _
		if win {
			b.PrintWithOptions(cfg.Color, cfg.Big, winLine) // ← было b.Print()
			name := cfg.NameX
			if current == "O" {
				name = cfg.NameO
			}
			fmt.Printf("%s wins!\n", name)
			if current == "X" {
				stats.WinsX++
			} else {
				stats.WinsO++
			}
			stats.Games++
			return
		}

		if b.IsFull() {
			b.PrintWithOptions(cfg.Color, cfg.Big, nil) // ← было b.Print()
			fmt.Println("Draw!")
			stats.Draws++
			stats.Games++
			return
		}

		if current == "X" {
			current = "O"
		} else {
			current = "X"
		}
	}
}

func getInput(b *board.Board, name string, size int) int {
	for {
		fmt.Printf("%s move: ", name)
		var cell int
		_, err := fmt.Scan(&cell)
		if err != nil {
			fmt.Printf("Error: enter a number 1-%d\n", size*size)
			var trash string
			fmt.Scan(&trash)
			continue
		}
		if cell < 1 || cell > size*size {
			fmt.Printf("Error: enter a number 1-%d\n", size*size)
			continue
		}
		if b.IsTaken(cell) {
			fmt.Printf("Error: cell %d is taken\n", cell)
			continue
		}
		return cell
	}
}

func printStats(stats Stats, verbose bool) {
	fmt.Println("=== Stats ===")
	fmt.Printf("Games: %d   X: %d   O: %d   Draws: %d\n",
		stats.Games, stats.WinsX, stats.WinsO, stats.Draws)

	if verbose && stats.Games > 0 {
		winRateX := stats.WinsX * 100 / stats.Games
		winRateO := stats.WinsO * 100 / stats.Games
		fmt.Printf("Moves this game: %d   Win rate — X: %d%%  O: %d%%\n",
			stats.MovesThis, winRateX, winRateO)
	}
}

func Exit(stats Stats) {
	printStats(stats, false)
	os.Exit(0)
}
