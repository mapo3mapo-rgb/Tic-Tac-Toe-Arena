package cli

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Mode    string
	Color   bool
	Big     bool
	First   string
	NameX   string
	NameO   string
	Size    int
	Verbose bool
}

func Parse() Config {
	cfg := Config{
		First: "X",
		Size:  3,
		NameX: "X",
		NameO: "O",
	}

	modeCount := 0

	for i := 1; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "--help", "-h":
			printUsage()
			os.Exit(0)
		case "--players":
			cfg.Mode = "players"
			modeCount++
		case "--ai":
			cfg.Mode = "ai"
			modeCount++
		case "--color":
			cfg.Color = true
		case "--big":
			cfg.Big = true
		case "--verbose":
			cfg.Verbose = true
		case "--first":
			i++
			cfg.First = os.Args[i]
		case "--name":
			i++
			parts := strings.Split(os.Args[i], ",")
			cfg.NameX = parts[0]
			cfg.NameO = parts[1]
		case "--size":
			i++
			n, err := strconv.Atoi(os.Args[i])
			if err != nil {
				fmt.Println("Error: --size must be a number")
				os.Exit(1)
			}
			cfg.Size = n
		default:
			fmt.Printf("Error: unknown flag %s\n", os.Args[i])
			printUsage()
			os.Exit(1)
		}
	}

	if modeCount != 1 {
		fmt.Println("Error: choose exactly one of --players or --ai")
		printUsage()
		os.Exit(1)
	}

	if cfg.First != "X" && cfg.First != "O" {
		fmt.Println("Error: --first must be X or O")
		printUsage()
		os.Exit(1)
	}

	if cfg.Size < 3 {
		fmt.Println("Error: --size must be an integer >= 3")
		printUsage()
		os.Exit(1)
	}

	if cfg.Mode == "ai" && cfg.Size != 3 {
		fmt.Println("Error: --ai and --size cannot be combined (AI is 3x3 only)")
		printUsage()
		os.Exit(1)
	}

	return cfg
}

func printUsage() {
	fmt.Println(`Usage: go run . (--players | --ai) [options]

Modes (exactly one required):
  --players        two human players take turns
  --ai             play against the computer (you are X)

Options:
  --color          enable colored output (default: plain)
  --big            render the board with large glyphs
  --verbose        show extended statistics
  --first X|O      who moves first (default: X)
  --name A,B       custom names: X=A, O=B
  --size N         board is N×N (default: 3)
  --help, -h       print this help and exit 0`)
}
