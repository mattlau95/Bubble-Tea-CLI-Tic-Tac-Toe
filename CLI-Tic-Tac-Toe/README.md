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

## 🛠️ Future Roadmap

- [ ] **Game Statistics:** Add a persistent session counter for Wins, Losses, and Draws.
- [ ] **Multiplayer:** Local hot-seat mode for two human players.
- [ ] **Animations:** Explore the `bubbles` package for simple timer or transition effects.

---
*Developed as a deep dive into idiomatic Go and modern TUI design.*