<div align="center">

# Checkers (Draughts) – Go CLI

A simple, terminal-based American Checkers (English Draughts) game written in Go.

</div>

## Overview

This project implements the core mechanics of an 8x8 checkers game with a text-based user interface. It supports:

- Turn-based play (Red vs Black)
- Standard diagonal movement and captures
- Multi-jump sequences entered in a single command (e.g. `C1E3G5`)
- Forced captures (cannot make a quiet move if a capture is available)
- Piece promotion to Queen (King) on reaching the far rank
- Win detection when one side has no remaining pieces

The code is intentionally compact and readable as a learning exercise in structuring a small Go CLI game.

## Installation

### Prerequisites

- Go 1.21+ (earlier versions likely work, but not tested)

### Clone & Build

```bash
git clone https://github.com/neilsmahajan/checkers.git
cd checkers
make build   # or: go build -o bin/checkers ./cmd/checkers
```

### Run

```bash
./bin/checkers
```

## How to Play

The board is labeled with rows `A`–`H` (top to bottom) and columns `1`–`8` (left to right).

Pieces:

- Red normal piece: 🔴
- Red queen: 🟥
- Black normal piece: ⚫️
- Black queen: ⬛️

### Input Format

Enter coordinates as a concatenated sequence of positions: `<StartRow><StartCol><EndRow><EndCol>[<NextEndRow><NextEndCol>...]`

Examples:

- Single move: `C3D4`
- Single capture: `C3E5`
- Multi-jump capture: `C3E5G7`
- Quit: `Q`

Inputs are case-insensitive; rows must be letters A–H and columns digits 1–8.

### Movement Rules

- Normal (non-queen) Red pieces move toward increasing row letters (down the printed board) diagonally.
- Normal Black pieces move toward decreasing row letters (up the printed board) diagonally.
- Queens (kings) move and capture both forward and backward.
- Captures jump two diagonals over an opposing piece into an empty square.
- Multiple captures in a single turn must be chained and entered in one command.
- If any capture is available for the current player, a non-capturing move is rejected.

### Promotion

A piece is promoted to a queen upon reaching the opponent's back rank (row `H` for Red, row `A` for Black).

### Win Condition

The game ends when one player has no remaining pieces. (Note: stalemate or no-move-available scenarios are not currently detected.)

## Project Structure

```
cmd/checkers      # Entry point (main package)
internal/board    # Core game logic: board state, movement, validation
internal/cli      # Command-line loop, input parsing, turn handling
bin/              # Build artifacts (ignored in VCS if added to .gitignore)
```

## Development

Common tasks:

```bash
make build   # Builds binary to bin/checkers
make run     # (If added) run the game directly
go vet ./...
go test ./...   # (No tests yet)
```

## Notable Implementation Details

- Input parsing batches an entire multi-jump sequence, validates it with `ValidateMoveSequence`, then executes atomically.
- Forced capture logic uses `PlayerHasCapture` before allowing a simple move.
- Board copies are used for dry-run validation of multi-jump sequences.
- Direction restrictions for non-queen pieces are enforced centrally.

## Current Limitations / Potential Improvements

1. No automated tests. Adding tests for movement, capture enforcement, promotion, multi-jump continuity, and win detection would harden the logic.
2. Turn switching after capture does not currently enforce "continue capturing with the same piece" if further captures are available; multi-jump only works if the user enters all segments up front. (Could auto-detect and require continuation.)
3. Stalemate / no-legal-move detection is not implemented; only piece exhaustion ends the game.
4. Error handling UX could be refined (e.g., clearer differentiation between parse vs. rule errors).
5. `SwitchTurn` TODO comment about not switching after an incomplete capture chain remains partially addressed logically by requiring the full chain in one input, but could be refactored.
6. Unicode piece symbols may not render consistently on all terminals; optionally provide ASCII fallback.
7. Promotion is immediate, but subsequent part of a multi-jump after promotion (rare in American rules) isn't re-evaluated for additional capture opportunities.
8. No CI pipeline or lint configuration file (like golangci-lint) yet.

## Contributing

Issues and pull requests are welcome. For larger changes, please open an issue first to discuss scope.

Suggested next steps if you want to extend the project:

- Add unit tests (table-driven tests for move legality)
- Implement forced continuation of capture chains
- Add stalemate detection (no legal moves)
- Introduce an AI opponent (minimax / alpha-beta)
- Add PGN-like move logging
- Provide a web or TUI (text UI) interface

## License

This project is licensed under the MIT License – see `LICENCE` for details.

## Author

Neil Mahajan  
Email: <neilsmahajan@gmail.com>  
Links: https://links.neilsmahajan.com

## Quick Demo (Example Session)

```
Checkers CLI
    1    2    3    4    5    6    7    8
  +----+----+----+----+----+----+----+----+
A | 🔴 |    | 🔴 |    | 🔴 |    | 🔴 |    |
  +----+----+----+----+----+----+----+----+
B |    | 🔴 |    | 🔴 |    | 🔴 |    | 🔴 |
  +----+----+----+----+----+----+----+----+
C | 🔴 |    | 🔴 |    | 🔴 |    | 🔴 |    |
  +----+----+----+----+----+----+----+----+
D |    |    |    |    |    |    |    |    |
  +----+----+----+----+----+----+----+----+
E |    |    |    |    |    |    |    |    |
  +----+----+----+----+----+----+----+----+
F |    | ⚫️ |    | ⚫️ |    | ⚫️ |    | ⚫️ |
  +----+----+----+----+----+----+----+----+
G | ⚫️ |    | ⚫️ |    | ⚫️ |    | ⚫️ |    |
  +----+----+----+----+----+----+----+----+
H |    | ⚫️ |    | ⚫️ |    | ⚫️ |    | ⚫️ |
  +----+----+----+----+----+----+----+----+
It's Red's Turn
Input start and end position in format <Start Row><Start Column><End Row><End Column> (e.g. C1D2)
To jump multiple pieces, chain positions (e.g. C1E3G5). Use Q to quit.
C1D2
Move from C1 to D2
```

## Acknowledgements

Built as a practice project in Go; inspired by classic board game mechanics.

---

Feel free to reach out if you have suggestions or feature ideas.
