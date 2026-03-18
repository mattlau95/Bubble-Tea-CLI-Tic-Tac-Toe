package main


import "fmt"

type Board [3][3]string

func CreateBoard() Board {
	board := Board{}
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			board[r][c] = " "	
		}
	}

	return board
}

func PrintBoard(board Board) {

	fmt.Println()
	fmt.Println("    1   2   3")
	fmt.Println()

	for r := 0; r < 3; r++ {

		rowLabel := string('A' + r)

		fmt.Printf("%s   %s | %s | %s \n",
			rowLabel,
			board[r][0],
			board[r][1],
			board[r][2],
		)

		if r < 2 {
			fmt.Println("   ---+---+---")
		}
	}

	fmt.Println()
}

func MakeMove(board *Board, row int, col int, player string) bool {

	if board[row][col] != " " {
		return false
	}

	board[row][col] = player
	return true
}

func CheckWinner(board Board) string {

	// rows
	for r := 0; r < 3; r++ {
		if board[r][0] != " " &&
			board[r][0] == board[r][1] &&
			board[r][1] == board[r][2] {
			return board[r][0]
		}
	}

	// columns
	for c := 0; c < 3; c++ {
		if board[0][c] != " " &&
			board[0][c] == board[1][c] &&
			board[1][c] == board[2][c] {
				return board[0][c]
		}
	}

	// diagonals
	if board[0][0] != " " &&  
		board[0][0] == board[1][1] &&
		board[1][1] == board[2][2] {
			return board[0][0]
			
	}
	
	
	if board[0][2] != " " && 
		board[0][2] == board[1][1] &&
		board[1][1] == board[2][0] {
			return board[0][2]

	}

	return ""

}

func IsBoardFull(board Board) bool {

	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			if board[r][c] == " " {
				return false
			}
		}
	}
	
	return true
}

func GetAvailableMoves(board Board) [][2]int {

	var moves [][2]int

	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			if board[r][c] == " " {
				moves = append(moves, [2]int{r, c})
			}
		}
	}

	return moves
}

func CopyBoard(board Board) Board {

	newBoard := board

	return newBoard
}