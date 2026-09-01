# Tic-Tac-Toe Arena

A terminal tic-tac-toe game with two play modes, a computer opponent, colored output, and large glyphs.

## How to run

```bash
go run . (--players | --ai) [options]
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
$ go run . --ai
1 | 2 | 3
---+---+---
4 | 5 | 6
---+---+---
7 | 8 | 9
X move: 5
...
O wins!


## Team

| Name | GitHub |
|---|---|
| Maxat Khiyuazha | @makhiy |
| David Abramov | @dabramov |
| Malika Kossymbayeva | @mkossy |
## Testing
All flags verified and tested by @khiyuazha-cell
