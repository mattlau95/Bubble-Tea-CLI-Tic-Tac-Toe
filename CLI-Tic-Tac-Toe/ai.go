package main

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomMove(board Board) (int, int) {
	delay(2)
	rand.Seed(time.Now().UnixNano())
	for {

		row := rand.Intn(3)
		col := rand.Intn(3)

		if board[row][col] == " " {
			return row, col
		}
	}
}

func MediumMove(board Board) (int, int) {

	delay(3)

	// 1. Try to win
	moves := GetAvailableMoves(board)

	for _, move := range moves {

		testBoard := CopyBoard(board)

		MakeMove(&testBoard, move[0], move[1], "O")

		if CheckWinner(testBoard) == "O" {
			return move[0], move[1]
		}
	}

	// 2. Block player win
	for _, move := range moves {

		testBoard := CopyBoard(board)

		MakeMove(&testBoard, move[0], move[1], "X")

		if CheckWinner(testBoard) == "X" {
			return move[0], move[1]
		}
	}

	// 3. Otherwise random
	return RandomMove(board)
}

func delay(seconds int) {
	fmt.Println("CPU is thinking...")
    
	// Cast seconds to time.Duration to allow multiplication
	time.Sleep(time.Duration(seconds) * time.Second)
    
	fmt.Println("Got it!")
}

func Minimax(board Board, depth int, isMaximizing bool, alpha int, beta int) int {

	winner := CheckWinner(board)

	if winner == "O" {
		return 10 - depth
	}

	if winner == "X" {
		return depth - 10
	}

	if IsBoardFull(board) {
		return 0
	}

	moves := GetAvailableMoves(board)

	if isMaximizing {

		bestScore := -1000

		for _, move := range moves {

			testBoard := CopyBoard(board)
			MakeMove(&testBoard, move[0], move[1], "O")

			score := Minimax(testBoard, depth+1, false, alpha, beta)

			if score > bestScore {
				bestScore = score
			}

			if bestScore > alpha {
				alpha = bestScore
			}

			if beta <= alpha {
				break // 🔥 prune
			}
		}

		return bestScore

	} else {

		bestScore := 1000

		for _, move := range moves {

			testBoard := CopyBoard(board)
			MakeMove(&testBoard, move[0], move[1], "X")

			score := Minimax(testBoard, depth+1, true, alpha, beta)

			if score < bestScore {
				bestScore = score
			}

			if bestScore < beta {
				beta = bestScore
			}

			if beta <= alpha {
				break // 🔥 prune
			}
		}

		return bestScore
	}
}

func BestMove(board Board) (int, int) {
	delay(3)


	bestScore := -1000
	bestRow := 0
	bestCol := 0

	moves := GetAvailableMoves(board)

	for _, move := range moves {

		testBoard := CopyBoard(board)
		MakeMove(&testBoard, move[0], move[1], "O")

		score := Minimax(testBoard, 0, false, -1000, 1000)

		if score > bestScore {
			bestScore = score
			bestRow = move[0]
			bestCol = move[1]
		}
	}

	return bestRow, bestCol
}