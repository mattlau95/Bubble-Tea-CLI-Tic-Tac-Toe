package main

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// --- Styles ---
var (
	xStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#5A96E3")).Bold(true)
	oStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#E35A5A")).Bold(true)
	cursorStyle = lipgloss.NewStyle().Background(lipgloss.Color("#3C3C3C")).Foreground(lipgloss.Color("#FFFFFF"))
	titleStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#00FF00")).Underline(true).Bold(true)
	msgStyle    = lipgloss.NewStyle().Italic(true).Foreground(lipgloss.Color("#AAAAAA"))
	
	// Stats Styling
	statsStyle = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder(), false, false, false, true).
			BorderForeground(lipgloss.Color("#3C3C3C")).
			PaddingLeft(2).
			MarginTop(1)

	// Transition Border for Game Over
	gameOverBox = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#FFD700")).
			Padding(1, 2).
			MarginTop(1)
)

// --- Custom Messages & Types ---
type tickMsg struct{}
type cpuMoveMsg struct{ row, col int }

type stats struct {
	wins   int
	losses int
	draws  int
}

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

	// Stats & Animation State
	stats      stats
	frame      int
	isThinking bool
}

func initialModel() model {
	return model{
		board:         CreateBoard(),
		currentPlayer: "X",
		message:       "Your move",
		state:         stateSelectDifficulty,
		stats:         stats{0, 0, 0},
	}
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Millisecond*250, func(t time.Time) tea.Msg {
		return tickMsg{}
	})
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Global Quit
	if key, ok := msg.(tea.KeyMsg); ok {
		if key.String() == "q" || key.String() == "ctrl+c" {
			return m, tea.Quit
		}
	}

	// 1. Difficulty Selection
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

	// 2. Gameplay Logic
	switch msg := msg.(type) {
	case tickMsg:
		if m.isThinking {
			m.frame++
			return m, tickCmd()
		}

	case cpuMoveMsg:
		m.isThinking = false
		MakeMove(&m.board, msg.row, msg.col, "O")
		winner := CheckWinner(m.board)
		if winner == "O" {
			m.message, m.gameOver = "CPU wins!", true
			m.stats.losses++
		} else if IsBoardFull(m.board) {
			m.message, m.gameOver = "Draw!", true
			m.stats.draws++
		} else {
			m.currentPlayer, m.message = "X", "Your move"
		}

	case tea.KeyMsg:
		if m.gameOver {
			if msg.String() == "r" {
				m.board = CreateBoard()
				m.gameOver = false
				m.currentPlayer = "X"
				m.message = "Rematch! Your move"
				return m, nil
			}
			return m, nil
		}

		switch msg.String() {
		case "up":    if m.cursorRow > 0 { m.cursorRow-- }
		case "down":  if m.cursorRow < 2 { m.cursorRow++ }
		case "left":  if m.cursorCol > 0 { m.cursorCol-- }
		case "right": if m.cursorCol < 2 { m.cursorCol++ }
		case "enter":
			if m.board[m.cursorRow][m.cursorCol] == " " && m.currentPlayer == "X" && !m.isThinking {
				MakeMove(&m.board, m.cursorRow, m.cursorCol, "X")
				
				winner := CheckWinner(m.board)
				if winner == "X" {
					m.message, m.gameOver = "You win!", true
					m.stats.wins++
					return m, nil
				}
				if IsBoardFull(m.board) {
					m.message, m.gameOver = "Draw!", true
					m.stats.draws++
					return m, nil
				}

				m.isThinking = true
				m.currentPlayer = "O"
				return m, tea.Batch(tickCmd(), cpuMoveCmd(m.board, m.difficulty))
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

	boardView := m.renderBoard()
	
	// Create Scoreboard
	scoreboard := statsStyle.Render(fmt.Sprintf(
		"Wins: %d | Losses: %d | Draws: %d",
		m.stats.wins, m.stats.losses, m.stats.draws,
	))

	if m.gameOver {
		return gameOverBox.Render(
			boardView + "\n" + 
			lipgloss.NewStyle().Foreground(lipgloss.Color("#FFD700")).Bold(true).Render(m.message) + "\n" +
			scoreboard + "\n\n" +
			"Press 'r' for Rematch | 'q' to Quit",
		)
	}

	var status string
	if m.isThinking {
		dots := ""
		for i := 0; i < (m.frame % 4); i++ { dots += "." }
		status = fmt.Sprintf("CPU is thinking%s", dots)
	} else {
		status = fmt.Sprintf("%s (%s)", m.message, diffToString(m.difficulty))
	}

	return boardView + "\n" + msgStyle.Render(status) + "\n" + scoreboard + "\n\nUse arrows + Enter | q to quit"
}

func (m model) renderBoard() string {
	s := "\n    1   2   3\n"
	for r := 0; r < 3; r++ {
		rowLabel := string(rune('A' + r))
		s += fmt.Sprintf(" %s ", rowLabel)
		for c := 0; c < 3; c++ {
			cellRaw := m.board[r][c]
			var cellView string

			switch cellRaw {
			case "X": cellView = xStyle.Render("X")
			case "O": cellView = oStyle.Render("O")
			default:  cellView = "."
			}

			if r == m.cursorRow && c == m.cursorCol && !m.gameOver && !m.isThinking {
				s += cursorStyle.Render(fmt.Sprintf("[%s]", cellView))
			} else {
				s += fmt.Sprintf(" %s ", cellView)
			}
			if c < 2 { s += "|" }
		}
		s += "\n"
		if r < 2 { s += "   ---+---+---\n" }
	}
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
		// Artificial thinking time to let the user see the animation
		time.Sleep(1200 * time.Millisecond) 
		var row, col int
		switch diff {
		case Easy: row, col = RandomMove(board)
		case Medium: row, col = MediumMove(board)
		case Hard: row, col = BestMove(board)
		}
		return cpuMoveMsg{row: row, col: col}
	}
}