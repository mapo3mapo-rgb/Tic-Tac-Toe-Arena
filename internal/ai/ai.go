package ai

import "github.com/mapo3mapo-rgb/Tic-Tac-Toe-Arena/internal/board"

func GetMove(b board.Board) int {
	size := b.Size
	cells := b.Cells

	// Правило 1 победить
	if move := tryComplete(b, "O"); move != -1 {
		return move
	}

	// Правило 2 заблокировать
	if move := tryComplete(b, "X"); move != -1 {
		return move
	}

	// Правило 3 центр
	center := (size * size) / 2
	if cells[center] == "" {
		return center + 1
	}

	// Правило 4 угол
	corners := []int{0, size - 1, size * (size - 1), size*size - 1}
	for _, i := range corners {
		if cells[i] == "" {
			return i + 1
		}
	}

	// Правило 5 сторона
	for i, cell := range cells {
		if cell == "" {
			return i + 1
		}
	}

	return -1
}

func tryComplete(b board.Board, mark string) int {
	for i, cell := range b.Cells {
		if cell != "" {
			continue
		}
		b.Cells[i] = mark
		win, _ := b.CheckWin(mark)
		b.Cells[i] = ""
		if win {
			return i + 1
		}
	}
	return -1
}
