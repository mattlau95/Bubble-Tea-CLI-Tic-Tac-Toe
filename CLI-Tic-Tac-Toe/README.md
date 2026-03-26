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

* The Elm Architecture (MVU):
  The application follows the Model-View-Update pattern. This ensures a unidirectional data flow where the View is a pure function of the Model. By separating the state from the rendering logic, the TUI remains bug-free and easy to extend with new features like "Rematch" or "Statistics."

* Concurrent AI & Non-Blocking I/O:
  Heavy AI computations (Minimax) are offloaded into Bubble Tea "Commands" (tea.Cmd). This leverages Go's goroutines to calculate the best move in the background, allowing the UI to stay responsive and play "Thinking..." animations at a consistent frame rate without blocking the main execution thread.

* Optimized Minimax with Alpha-Beta Pruning:
  The "Hard" AI uses a recursive search to evaluate all possible board outcomes. To prevent performance lag, Alpha-Beta pruning is implemented to "cut off" branches in the search tree that cannot possibly influence the final decision, drastically reducing the number of nodes evaluated.

* Declarative CSS-like Styling:
  Using the Lip Gloss library, UI components are styled using a functional, declarative approach. This allows for complex layouts, borders, and ANSI color highlights (X in Blue, O in Red) that adapt gracefully to different terminal themes and sizes.

* Real-Time State Management:
  A custom tick-based system triggers periodic messages to the update loop. This facilitates real-time UI elements, such as the animated loading dots during the CPU's turn, making the terminal application feel fluid and modern.

---
Developed as a deep dive into idiomatic Go and modern TUI design.