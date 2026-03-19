# Tic-Tac-Toe TUI (Go)

An interactive, event-driven Terminal User Interface (TUI) for Tic-Tac-Toe. Built with Go and the Bubble Tea framework, featuring a recursive Minimax AI with Alpha-Beta pruning.

This project demonstrates a shift from procedural CLI logic to a functional, state-based architecture (Model-View-Update), providing a much smoother and more modern terminal experience.

## Features

* Interactive TUI: Navigate the board using Arrow Keys and select your move with Enter.
* Async AI Thinking: The UI remains responsive with a "CPU thinking..." status while the Minimax algorithm runs in the background.
* Three Difficulty Levels:
    * Easy: CPU makes entirely random moves.
    * Medium: CPU looks one move ahead to prioritize winning or blocking the player.
    * Hard: Uses a Minimax Algorithm with Alpha-Beta Pruning for unbeatable, optimal play.
* Clean Visuals: Labeled rows (A-C) and columns (1-3) with active cursor highlighting.

## Example UI

     1   2   3
 A  [X]| . | O 
    ---+---+---
 B   . | O | . 
    ---+---+---
 C   . | . | X 

CPU thinking...

Use arrow keys + Enter | q to quit

## How to Run

1. Prerequisites: Ensure you have Go 1.18 or higher installed.
2. Initialize & Install Dependencies: run 'go mod tidy'
3. Run the Game: run 'go run .'

## Project Architecture

The project follows a modular design to ensure a strict separation of concerns:

- main.go: Entry point; initializes the Bubble Tea program and terminal loop.
- ui.go: The TUI Controller. Manages state transitions, keyboard input, and rendering.
- ai.go: The AI Engine. Contains Random, Heuristic (Medium), and Minimax (Hard) logic.
- board.go: The Data Model. Handles the 3x3 board state and win/draw detection logic.
- game.go: Shared Definitions. Contains difficulty constants and coordinate parsing.

## Technical Highlights

* Minimax Optimization: Implemented Alpha-Beta pruning to significantly reduce the search space, ensuring the "Hard" mode is fast and impossible to beat.
* Elm Architecture: Utilizes the Bubble Tea framework's Model-View-Update pattern for predictable state management.
* Non-Blocking I/O: Leverages Go's concurrency model via tea.Cmd to perform AI calculations without freezing the user interface.

## Future Roadmap

- [ ] Lip Gloss Integration: Add colors (Blue for X, Red for O) and stylish borders.
- [ ] Score Tracking: Add a persistent session counter for Wins, Losses, and Draws.
- [ ] Multiplayer: Local hot-seat mode for two human players.

---
Developed as a deep dive into idiomatic Go and TUI design.