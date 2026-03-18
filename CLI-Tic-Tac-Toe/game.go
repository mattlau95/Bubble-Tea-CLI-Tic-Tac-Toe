package main

import (
	"fmt"
	"strings"
)

type Difficulty int

const (
	Easy Difficulty = iota
	Medium
	Hard
)

func ParseMove(input string) (int, int, bool) {

	if len(input) != 2 {
		return 0, 0, false
	}

	row := int(input[0] - 'A')
	col := int(input[1] - '1')

	if row < 0 || row > 2 || col < 0 || col > 2 {
		return 0, 0, false
	}

	return row, col, true
}

func PlayGame() {

	board := CreateBoard()
	currentPlayer := "X"

	difficulty := SelectDifficulty()

	for {

		PrintBoard(board)

		var row int
		var col int

		if currentPlayer == "X" {

			var move string

			fmt.Print("Enter move (example B2): ")
			fmt.Scan(&move)
			move = strings.ToUpper(move)

			var ok bool
			row, col, ok = ParseMove(move)

			if !ok {
				fmt.Println("Invalid input.")
				continue
			}

			if !MakeMove(&board, row, col, currentPlayer) {
				fmt.Println("That space is taken.")
				continue
			}

		} else {

			switch difficulty {
			case Easy:
				row, col = RandomMove(board)
			case Medium:
				row, col = MediumMove(board)
			case Hard:
				row, col = BestMove(board)
			}

			fmt.Println("CPU plays", FormatMove(row, col))

			MakeMove(&board, row, col, currentPlayer)
		}

		winner := CheckWinner(board)

		if winner != "" {
			PrintBoard(board)
			fmt.Println("Player", winner, "wins!")
			break
		}

		if IsBoardFull(board) {
			PrintBoard(board)
			fmt.Println("It's a draw!")
			break
		}

		if currentPlayer == "X" {
			currentPlayer = "O"
		} else {
			currentPlayer = "X"
		}
	}
}

func toBoardCoord(row, col int) string {
    // Convert col index to a letter (0 -> A, 1 -> B, etc.)
    // 'A' has an ASCII value of 65.
    columnLetter := string(rune('A' + col))
    
    // Convert row index to 1-based number (0 -> 1, 1 -> 2, etc.)
    rowNumber := row + 1
    
    // Format into a single string
    return fmt.Sprintf("%s%d", columnLetter, rowNumber)
}

func SelectDifficulty() Difficulty {

	var choice int

	fmt.Println("Select Difficulty:")
	fmt.Println("1 - Easy")
	fmt.Println("2 - Medium")
	fmt.Println("3 - Hard")
	fmt.Print("Enter choice: ")

	fmt.Scan(&choice)

	switch choice {
	case 1:
		return Easy
	case 2:
		return Medium
	case 3:
		return Hard
	default:
		fmt.Println("Invalid choice, defaulting to Medium.")
		return Medium
	}
}

func FormatMove(row, col int) string {
	return fmt.Sprintf("%c%d", 'A'+row, col+1)
}