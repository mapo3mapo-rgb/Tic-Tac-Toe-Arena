package board

import (
	"fmt"
	"strings"
)

const (
	colorRed   = "\033[91m"
	colorBlue  = "\033[94m"
	colorGreen = "\033[32;1m"
	colorDim   = "\033[2m"
	colorReset = "\033[0m"
)

type Board struct {
	Cells []string
	Size  int
}

func New(size int) Board {
	cells := make([]string, size*size)
	return Board{Cells: cells, Size: size}
}

func (b *Board) Reset() {
	for i := range b.Cells {
		b.Cells[i] = ""
	}
}

func (b *Board) Place(cell int, mark string) {
	b.Cells[cell-1] = mark
}

func (b *Board) IsTaken(cell int) bool {
	return b.Cells[cell-1] != ""
}

func cellContent(b Board, i int, color bool, winning []int) string {
	cell := b.Cells[i]

	isWinning := false
	for _, w := range winning {
		if w == i {
			isWinning = true
			break
		}
	}

	if cell == "" {
		num := fmt.Sprintf("%d", i+1)
		if color {
			return colorDim + num + colorReset
		}
		return num
	}

	if color {
		if isWinning {
			return colorGreen + cell + colorReset
		}
		if cell == "X" {
			return colorRed + cell + colorReset
		}
		return colorBlue + cell + colorReset
	}
	return cell
}

func (b Board) Print() {
	b.PrintWithOptions(false, false, nil)
}

func (b Board) PrintColor() {
	b.PrintWithOptions(true, false, nil)
}

func (b Board) PrintWithOptions(color bool, big bool, winning []int) {
	size := b.Size
	if big {
		printBig(b, color, winning)
		return
	}
	for row := 0; row < size; row++ {
		line := ""
		for col := 0; col < size; col++ {
			i := row*size + col
			line += fmt.Sprintf(" %s ", cellContent(b, i, color, winning))
			if col < size-1 {
				line += "|"
			}
		}
		fmt.Println(line)
		if row < size-1 {
			fmt.Println(strings.Repeat("---+", size-1) + "---")
		}
	}
}

func (b Board) CheckWin(mark string) (bool, []int) {
	size := b.Size
	lines := [][]int{}

	for i := 0; i < size; i++ {
		row := []int{}
		col := []int{}
		for j := 0; j < size; j++ {
			row = append(row, i*size+j)
			col = append(col, j*size+i)
		}
		lines = append(lines, row, col)
	}

	diag1 := []int{}
	diag2 := []int{}
	for i := 0; i < size; i++ {
		diag1 = append(diag1, i*size+i)
		diag2 = append(diag2, i*size+(size-1-i))
	}
	lines = append(lines, diag1, diag2)

	for _, line := range lines {
		win := true
		for _, idx := range line {
			if b.Cells[idx] != mark {
				win = false
				break
			}
		}
		if win {
			return true, line
		}
	}
	return false, nil
}

func (b Board) IsFull() bool {
	for _, cell := range b.Cells {
		if cell == "" {
			return false
		}
	}
	return true
}

func printBig(b Board, color bool, winning []int) {
	size := b.Size
	xGlyph := []string{"X   X", "  X  ", "X   X"}
	oGlyph := []string{" OOO ", "O   O", " OOO "}

	for row := 0; row < size; row++ {
		for line := 0; line < 3; line++ {
			out := ""
			for col := 0; col < size; col++ {
				i := row*size + col
				cell := b.Cells[i]

				isWinning := false
				for _, w := range winning {
					if w == i {
						isWinning = true
						break
					}
				}

				var content string
				if cell == "X" {
					content = xGlyph[line]
					if color {
						if isWinning {
							content = colorGreen + content + colorReset
						} else {
							content = colorRed + content + colorReset
						}
					}
				} else if cell == "O" {
					content = oGlyph[line]
					if color {
						if isWinning {
							content = colorGreen + content + colorReset
						} else {
							content = colorBlue + content + colorReset
						}
					}
				} else {
					if line == 1 {
						content = fmt.Sprintf("  %d  ", i+1)
					} else {
						content = "     "
					}
				}

				out += content
				if col < size-1 {
					out += "|"
				}
			}
			fmt.Println(out)
		}
		if row < size-1 {
			fmt.Println(strings.Repeat("-----+", size-1) + "-----")
		}
	}
}
