package main

import (
	"math/rand"
	"time"
)

func RandomMove(board Board) (int, int) {

	rand.Seed(time.Now().UnixNano())

	for {

		row := rand.Intn(3)
		col := rand.Intn(3)

		if board[row][col] == " " {
			return row, col
		}
	}
}

