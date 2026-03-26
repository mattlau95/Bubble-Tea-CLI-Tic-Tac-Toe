package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Define Styles
var (
	xStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#5A96E3")).Bold(true) // Soft Blue
	oStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#E35A5A")).Bold(true) // Soft Red
	cursorStyle = lipgloss.NewStyle().Background(lipgloss.Color("#3C3C3C")).Foreground(lipgloss.Color("#FFFFFF"))
	titleStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#00FF00")).Underline(true).Bold(true)
	msgStyle    = lipgloss.NewStyle().Italic(true).Foreground(lipgloss.Color("#AAAAAA"))
)

type sessionState int

const (
	stateSelectDifficulty sessionState = iota
	statePlaying
)

type model struct {
	board         Board
	currentPlayer string
	cursorRow     int
	cursorCol     int
	message       string
	gameOver      bool
	difficulty    Difficulty
	state         sessionState
}

func initialModel() model {
	return model{
		board:         CreateBoard(),
		currentPlayer: "X",
		cursorRow:     0,
		cursorCol:     0,
		message:       "Your move",
		gameOver:      false,
		state:         stateSelectDifficulty,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if key, ok := msg.(tea.KeyMsg); ok {
		if key.String() == "q" || key.String() == "ctrl+c" {
			return m, tea.Quit
		}
	}

	if m.state == stateSelectDifficulty {
		if msg, ok := msg.(tea.KeyMsg); ok {
			switch msg.String() {
			case "1": m.difficulty, m.state = Easy, statePlaying
			case "2": m.difficulty, m.state = Medium, statePlaying
			case "3": m.difficulty, m.state = Hard, statePlaying
			}
		}
		return m, nil
	}

	if m.gameOver {
		return m, tea.Quit
	}

	switch msg := msg.(type) {
	case cpuMoveMsg:
		MakeMove(&m.board, msg.row, msg.col, "O")
		winner := CheckWinner(m.board)
		if winner != "" {
			m.message, m.gameOver = "CPU wins!", true
		} else if IsBoardFull(m.board) {
			m.message, m.gameOver = "Draw!", true
		} else {
			m.currentPlayer, m.message = "X", "Your move"
		}
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "up": if m.cursorRow > 0 { m.cursorRow-- }
		case "down": if m.cursorRow < 2 { m.cursorRow++ }
		case "left": if m.cursorCol > 0 { m.cursorCol-- }
		case "right": if m.cursorCol < 2 { m.cursorCol++ }
		case "enter":
			if m.board[m.cursorRow][m.cursorCol] == " " && m.currentPlayer == "X" {
				MakeMove(&m.board, m.cursorRow, m.cursorCol, "X")
				if CheckWinner(m.board) != "" {
					m.message, m.gameOver = "You win!", true
					return m, nil
				}
				if IsBoardFull(m.board) {
					m.message, m.gameOver = "Draw!", true
					return m, nil
				}
				m.currentPlayer, m.message = "O", "CPU thinking..."
				return m, cpuMoveCmd(m.board, m.difficulty)
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	if m.state == stateSelectDifficulty {
		return "\n  " + titleStyle.Render("TIC-TAC-TOE") + "\n\n" +
			"  Select Difficulty:\n" +
			"  1. Easy\n" +
			"  2. Medium\n" +
			"  3. Hard\n\n" +
			msgStyle.Render("Press a number to start...")
	}

	s := "\n    1   2   3\n"
	for r := 0; r < 3; r++ {
		rowLabel := string(rune('A' + r))
		s += fmt.Sprintf(" %s ", rowLabel)
		for c := 0; c < 3; c++ {
			cellRaw := m.board[r][c]
			var cellView string

			// Apply Player Styles
			switch cellRaw {
			case "X": cellView = xStyle.Render("X")
			case "O": cellView = oStyle.Render("O")
			default:  cellView = "."
			}

			// Apply Cursor Style
			if r == m.cursorRow && c == m.cursorCol {
				s += cursorStyle.Render(fmt.Sprintf("[%s]", cellView))
			} else {
				s += fmt.Sprintf(" %s ", cellView)
			}

			if c < 2 { s += "|" }
		}
		s += "\n"
		if r < 2 { s += "   ---+---+---\n" }
	}

	status := fmt.Sprintf("\n%s (%s)\n", m.message, diffToString(m.difficulty))
	s += msgStyle.Render(status)
	s += "\nUse arrow keys + Enter | q to quit\n"
	return s
}

func diffToString(d Difficulty) string {
	switch d {
	case Easy: return "Easy"
	case Medium: return "Medium"
	case Hard: return "Hard"
	default: return ""
	}
}

func cpuMoveCmd(board Board, diff Difficulty) tea.Cmd {
	return func() tea.Msg {
		var row, col int
		switch diff {
		case Easy: row, col = RandomMove(board)
		case Medium: row, col = MediumMove(board)
		case Hard: row, col = BestMove(board)
		}
		return cpuMoveMsg{row: row, col: col}
	}
}

type cpuMoveMsg struct{ row, col int }