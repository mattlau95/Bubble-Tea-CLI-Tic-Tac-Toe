package main

import (
	"fmt"
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

		var move string

		fmt.Print("Enter move (example B2): ")
		fmt.Scan(&move)

		row, col, ok := ParseMove(move)

		if !ok {
			fmt.Println("Invalid input. Use format like A1 or C3.")
			continue
		}

		validMove := MakeMove(&board, row, col, currentPlayer)

		if !validMove {
			fmt.Println("Invalid move! Try again.")
			continue
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