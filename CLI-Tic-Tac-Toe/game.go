package main

import (
	"fmt"
	"strings"
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

			row, col = RandomMove(board)

			fmt.Println("CPU chooses:", toBoardCoord(row, col))

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