package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

type cpuMoveMsg struct {
    row int
    col int
}
type model struct {
	board         Board
	currentPlayer string
	cursorRow     int
	cursorCol     int
	message       string
	gameOver      bool
}

func initialModel() model {
	return model{
		board:         CreateBoard(),
		currentPlayer: "X",
		cursorRow:     0,
		cursorCol:     0,
		message:       "Your move",
		gameOver:      false,
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	if m.gameOver {
		return m, tea.Quit
	}

	switch msg := msg.(type) {

	// ✅ HANDLE CPU RESULT HERE
	case cpuMoveMsg:

		MakeMove(&m.board, msg.row, msg.col, "O")

		winner := CheckWinner(m.board)
		if winner != "" {
			m.message = "CPU wins!"
			m.gameOver = true
			return m, nil
		}

		if IsBoardFull(m.board) {
			m.message = "Draw!"
			m.gameOver = true
			return m, nil
		}

		m.currentPlayer = "X"
		m.message = "Your move"
		return m, nil

	// ✅ HANDLE USER INPUT
	case tea.KeyMsg:

		switch msg.String() {

		case "q", "ctrl+c":
			return m, tea.Quit

		case "up":
			if m.cursorRow > 0 {
				m.cursorRow--
			}

		case "down":
			if m.cursorRow < 2 {
				m.cursorRow++
			}

		case "left":
			if m.cursorCol > 0 {
				m.cursorCol--
			}

		case "right":
			if m.cursorCol < 2 {
				m.cursorCol++
			}

		case "enter":

			if m.board[m.cursorRow][m.cursorCol] != " " {
				m.message = "Spot taken!"
				return m, nil
			}

			// Player move
			MakeMove(&m.board, m.cursorRow, m.cursorCol, "X")

			winner := CheckWinner(m.board)
			if winner != "" {
				m.message = "You win!"
				m.gameOver = true
				return m, nil
			}

			if IsBoardFull(m.board) {
				m.message = "Draw!"
				m.gameOver = true
				return m, nil
			}

			// ✅ Trigger CPU async
			m.currentPlayer = "O"
			m.message = "CPU thinking..."
			return m, cpuMoveCmd(m.board)
		}
	}

	return m, nil
}

func (m model) View() string {

	s := "\n"

	// Header
	s += "    1   2   3\n"

	for r := 0; r < 3; r++ {

		rowLabel := string(rune('A' + r))
		s += fmt.Sprintf(" %s ", rowLabel)

		for c := 0; c < 3; c++ {

			cell := m.board[r][c]

			if cell == " " {
				cell = "."
			}

			// Cursor highlight
			if r == m.cursorRow && c == m.cursorCol {
				s += fmt.Sprintf("[%s]", cell)
			} else {
				s += fmt.Sprintf(" %s ", cell)
			}

			if c < 2 {
				s += "|"
			}
		}

		s += "\n"

		if r < 2 {
			s += "   ---+---+---\n"
		}
	}

	s += "\n" + m.message + "\n"
	s += "\nUse arrow keys + Enter | q to quit\n"

	return s
}

func (m model) Init() tea.Cmd {
	return nil
}

func cpuMoveCmd(board Board) tea.Cmd {
	return func() tea.Msg {

		row, col := BestMove(board)

		return cpuMoveMsg{row: row, col: col}
	}
}