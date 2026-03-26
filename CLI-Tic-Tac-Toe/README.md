# Tic-Tac-Toe TUI (Go)

An interactive, event-driven Terminal User Interface (TUI) for Tic-Tac-Toe. Built with **Go**, the **Bubble Tea** framework, and **Lip Gloss** for styling. This project features a recursive Minimax AI with Alpha-Beta pruning, ensuring an unbeatable "Hard" mode.

This project demonstrates a shift from procedural CLI logic to a functional, state-based architecture (Model-View-Update), providing a polished and responsive terminal experience.

## 🕹️ Features

* **Interactive TUI:** Navigate the board using **Arrow Keys** and select your move with **Enter**.
* **Difficulty Selection:** A dedicated start screen to choose between Easy, Medium, and Hard modes.
* **Async AI Thinking:** The UI remains responsive with a "CPU thinking..." status while the Minimax algorithm runs in the background via `tea.Cmd`.
* **Rich Styling:** Powered by **Lip Gloss**, featuring:
    * **X** and **O** color coding (Blue/Red).
    * Active cursor highlighting with background tints.
    * Bold titles and italicized status messages.

## 📸 Example UI

     TIC-TAC-TOE

     1   2   3
 A  [X]| . | O 
    ---+---+---
 B   . | O | . 
    ---+---+---
 C   . | . | X 

CPU thinking... (Hard)

Use arrow keys + Enter | q to quit

## 🚀 How to Run

1.  **Prerequisites:** Ensure you have Go 1.18 or higher installed.
2.  **Initialize & Install Dependencies:**
    ```bash
    go mod tidy
    ```
3.  **Run the Game:**
    ```bash
    go run .
    ```

## 🏗️ Project Architecture

The project is designed with a strict separation of concerns, ensuring the AI logic, board state, and user interface remain modular and maintainable.

| Component | File | Responsibility |
| :--- | :--- | :--- |
| **Entry Point** | `main.go` | Initializes the Bubble Tea program and terminal loop. |
| **Controller** | `ui.go` | Manages state transitions, keyboard input, and TUI rendering. |
| **AI Engine** | `ai.go` | Contains Random, Heuristic (Medium), and Minimax (Hard) logic. |
| **Data Model** | `board.go` | Handles the 3x3 board state and win/draw detection. |
| **Definitions** | `game.go` | Stores shared difficulty constants and coordinate parsing. |

## 🧠 Technical Highlights

* **The Elm Architecture:** Utilizes Bubble Tea’s **Model-View-Update** pattern for predictable state management.
* **Lip Gloss Styling:** Uses a declarative styling approach to wrap strings in ANSI escape codes, ensuring cross-terminal compatibility for colors and layouts.
* **Minimax Optimization:** Implemented Alpha-Beta pruning to significantly reduce the search space, making the AI both fast and unbeatable.
* **Non-Blocking I/O:** Leverages Go's concurrency model to perform heavy AI calculations without freezing the main UI thread.

Styling: 
Colors!: X is now blue and O is red.

The Cursor: Instead of plain brackets [ ], the entire cell now highlights with a background color using cursorStyle.

Visual Hierarchy: Added a titleStyle for the menu and a msgStyle for status updates to make the text easier to scan.

Clean Code: Used Render() to wrap strings in ANSI escape codes, ensuring they display correctly in any modern terminal.


Animations and Transitions
The Pulse: The tickCmd sends a message every 250ms. The Update function catches it and increments m.frame, which the View uses to draw "CPU thinking.", "CPU thinking..", etc.

The Transition: When m.gameOver hits, the View switches to a gameOverBox. This puts a gold-rounded border around the final state of the board, visually separating the "end screen" from the gameplay.

Artificial Latency: I added a time.Sleep(1200 * time.Millisecond) inside cpuMoveCmd. Since it's inside a tea.Cmd, it won't freeze your UI, but it gives the user enough time to actually see the "Thinking..." animation.

## 🛠️ Future Roadmap

- [ ] **Game Statistics:** Add a persistent session counter for Wins, Losses, and Draws.
- [ ] **Multiplayer:** Local hot-seat mode for two human players.
- [ ] **Animations:** Explore the `bubbles` package for simple timer or transition effects.

---
*Developed as a deep dive into idiomatic Go and modern TUI design.*