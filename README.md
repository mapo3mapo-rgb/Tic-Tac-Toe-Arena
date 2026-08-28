# Tic-Tac-Toe Arena

A terminal tic-tac-toe game with two play modes, a computer opponent, colored output, and large glyphs.

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
| `--color` | Enable colored output |
| `--big` | Render the board with large glyphs |
| `--verbose` | Show extended statistics |
| `--first X\|O` | Who moves first (default: X) |
| `--name A,B` | Custom names: X=A, O=B |
| `--size N` | Board is N×N (default: 3) |
| `--help, -h` | Print help and exit |

## Rules

- Players take turns entering a cell number (1–9)
- First to complete a full row, column, or diagonal wins
- If the board fills with no winner — Draw
- AI applies a rule ladder: Win → Block → Center → Corner → Side

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
