# CLI Tic-Tac-Toe (Go)

A terminal-based Tic-Tac-Toe game built in Go with a focus on clean architecture, intuitive CLI UX, and progressively smarter AI opponents.

This project is part of my effort to strengthen my backend engineering and systems programming skills while designing polished terminal interfaces.

## Features

* Playable Tic-Tac-Toe game in the terminal
* Labeled board for intuitive input (A1, B2, C3)
* Input validation and error handling
* Turn-based player system
* Win and draw detection
* Clean modular architecture

## Example Board

```
    1   2   3
A   X | O |  
   ---+---+---
B     | X |  
   ---+---+---
C   O |   | X
```

Players enter moves using coordinates:

```
A1
B2
C3
```

## How to Run

Clone the repository and run:

```
go run .
```

## How to Play

1. The board is labeled with rows (A–C) and columns (1–3)
2. Players take turns entering coordinates
3. First player to align three marks wins
4. If the board fills without a winner, the game is a draw

## Project Structure

```
CLI-Tic-Tac-Toe/
│
├── main.go      # program entry point
├── board.go     # board state and operations
├── game.go      # game loop and player interaction
└── ai.go        # CPU player logic (coming soon)
```

## Future Improvements

* CPU opponent with multiple difficulty levels
* Minimax AI (unbeatable mode)
* Terminal UI enhancements
* Game statistics tracking
* Watch CPU vs CPU mode

## Learning Goals

This project focuses on:

* Writing idiomatic Go
* Designing clean CLI interfaces
* Implementing classic algorithms (Minimax)
* Structuring maintainable software projects
