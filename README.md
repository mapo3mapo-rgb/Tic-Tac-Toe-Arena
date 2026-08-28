# Tic-Tac-Toe Arena

A terminal tic-tac-toe game with two play modes, a computer opponent, colored output, and large glyphs. This is the final raid project that brings together everything from the course: structs, slices, strings, maps, errors, os.Args, and clean function decomposition.

## How to run

```bash
go run main.go (--players | --ai) [options]
```

### Modes (exactly one required)
| Flag | Description |
|---|---|
| `--players` | Two human players take turns |
| `--ai` | Play against the computer (you are X) |

### Options
| Flag | Description |
|---|---|
| `--color` | Enable colored output (default: plain) |
| `--big` | Render the board with large glyphs |
| `--verbose` | Show extended statistics |
| `--first X\|O` | Who moves first (default: X) |
| `--name A,B` | Custom names: X=A, O=B |
| `--size N` | Board is N×N (default: 3) |
| `--help, -h` | Print help and exit |

## Rules

### Two Modes
1. **`--players`** — Two human players take turns on the same terminal
2. **`--ai`** — You play as X against the computer (O). Use `--first O` to let the computer move first

### Win and Draw Detection
- **Win**: A complete row, column, or diagonal filled with the same mark
- **Draw**: The board is full with no winner
- On a win, the winning line is highlighted and the winner is announced
- On a draw, "Draw!" is announced

### How The AI Works
The AI is rule-based — no recursion, no minimax. On its turn, it applies this ladder and plays the first rule that matches:

1. **Win** — If O can complete a line this move, play it
2. **Block** — Else if X could complete a line next move, take that cell
3. **Center** — Else if the center (cell 5) is free, take it
4. **Corner** — Else take the first free corner, in order 1, 3, 7, 9
5. **Side** — Else take the first free side, in order 2, 4, 6, 8

The AI is fully deterministic — when several cells satisfy a rule, it always picks the one named first in the fixed order above.

## Example

```bash
$ go run . --ai
 1 | 2 | 3
---+---+---
 4 | 5 | 6
---+---+---
 7 | 8 | 9
X move: 5
 1 | 2 | 3
---+---+---
 4 | X | 6
---+---+---
 7 | 8 | 9
O plays 1
 O | 2 | 3
---+---+---
 4 | X | 6
---+---+---
 7 | 8 | 9
X move: 9
 O | 2 | 3
---+---+---
 4 | X | 6
---+---+---
 7 | 8 | X
O plays 3
 O | 2 | O
---+---+---
 4 | X | 6
---+---+---
 7 | 8 | X
X move: 7
 O | 2 | O
---+---+---
 4 | X | 6
---+---+---
 X | 8 | X
O plays 2
 O | O | O
---+---+---
 4 | X | 6
---+---+---
 X | 8 | X
O wins!
=== Stats ===
Games: 1   X: 0   O: 1   Draws: 0
Play again? (y/n): n
```

## Team

| Name | Pre-Piscine |
|---|---|
| Maxat Khiyuazha | @makhiy |
| David Abramov | @dabramov |
| Malika Kossymbayeva | @mkossy |

---
