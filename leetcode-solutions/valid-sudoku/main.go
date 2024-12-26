package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	// we have 9 rows, 9 cols and 9 squares
	// each of them can have only 9 values from 1 to 9
	// we can represent them all as fixed matrix of id of something (row, col, sq) plus number
	// with a bool value if this number is present here or not
	var rows, cols, squares [9][9]bool

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			ch := board[r][c]

			// dot mean empty value, so we don't check it
			if ch == '.' {
				continue
			}

			// based on the constraints we can have only 1-9
			// then we can safely convert them to int (basically index of a number)
			idx := ch - '1'

			// how to get the square index?
			// row / 3 and col / 3  - are going to give us values 0-2,
			// 0 | 1 | 2
			// --+---+--
			// 3 | 4 | 5
			// --+---+--
			// 6 | 7 | 8
			// looks like can do: (row / 3) * 3 + col / 3
			squareIdx := (r/3)*3 + c/3

			// check if we've seen this row, col or square
			if rows[r][idx] || cols[c][idx] || squares[squareIdx][idx] {
				return false
			}

			// store what we've seen
			rows[r][idx], cols[c][idx], squares[squareIdx][idx] = true, true, true
		}
	}

	return true
}

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	out := isValidSudoku(board)
	fmt.Printf("The given board validity check: %t\n", out)
}
