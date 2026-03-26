# Tic-Tac-Toe TUI (Go)

An interactive, event-driven Terminal User Interface (TUI) for Tic-Tac-Toe. Built with Go, the Bubble Tea framework, and Lip Gloss for styling. This project features a recursive Minimax AI with Alpha-Beta pruning, ensuring an unbeatable "Hard" mode.

This version moves beyond a simple CLI by implementing a full state machine, real-time animations, and persistent session statistics.

## 🕹️ Features

* Interactive Navigation: Use Arrow Keys to move the cursor and Enter to place your mark.
* Customizable Turn Order: Press [f] on the menu to toggle between Player (X) or CPU (O) going first.
* Three AI Difficulty Levels:
    - Easy: Random move selection.
    - Medium: Heuristic-based (prioritizes immediate wins/blocks).
    - Hard: Unbeatable Minimax algorithm with Alpha-Beta pruning.
* Live Animations: A tick-based "Thinking..." animation for the CPU to mimic human reaction time.
* Session Statistics: Persistent tracking of Wins, Losses, and Draws during your session.
* Seamless Rematch: Press [r] after a game ends to instantly clear the board and play again without restarting.

## 📸 Example UI

     TIC-TAC-TOE

     1   2   3
 A   X | . | O 
    ---+---+---
 B   . |[O]| . 
    ---+---+---
 C   . | . | X 

CPU is thinking... (Hard)

Wins: 2 | Losses: 1 | Draws: 4

Use arrows + Enter | q to quit

## 🚀 How to Run

1. Prerequisites: Ensure you have Go 1.18+ installed.
2. Initialize & Install Dependencies: Run 'go mod tidy'
3. Run the Game: Run 'go run .'

## 🏗️ Project Architecture

| Component | File | Responsibility |
| :--- | :--- | :--- |
| Entry Point | main.go | Initializes the Bubble Tea program loop. |
| Controller | ui.go | Manages TUI state, Lip Gloss styles, and animations. |
| AI Engine | ai.go | Contains the Minimax and Random move logic. |
| Data Model | board.go | Handles the 3x3 board state and win detection. |
| Definitions | game.go | Shared constants and difficulty enums. |

## 🧠 Technical Highlights

* Model-View-Update (MVU): Leveraging the Elm Architecture via Bubble Tea for predictable UI state transitions.
* Non-Blocking I/O: Using tea.Cmd and tea.Tick to handle background AI calculations and real-time UI updates simultaneously.
* Declarative Styling: Using Lip Gloss to define reusable UI components with ANSI colors and borders.
* Algorithm Efficiency: Minimax with Alpha-Beta pruning ensures the CPU plays optimally without high latency.

---
Developed as a deep dive into idiomatic Go and modern TUI design.