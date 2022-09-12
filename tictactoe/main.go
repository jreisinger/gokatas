// Tictactoe creates a a tic-tac-toe board as a slice of string slices. Adapted
// from tour.golang.org/moretypes/14.
//
// Level: beginner
// Topics: types, slice of slices
package main

import (
	"fmt"
	"strings"
)

func main() {
	board := [][]string{
		// []string{" ", " ", " "},
		{" ", " ", " "},
		{" ", " ", " "},
		{" ", " ", " "},
	}

	// The players take turns and play like this:
	// x   o
	//   x
	// o   x
	board[0][0] = "x"
	board[0][2] = "o"
	board[1][1] = "x"
	board[2][0] = "o"
	board[2][2] = "x"

	// for i := 0; i < len(board); i++ {
	for i := range board {
		fmt.Println(strings.Join(board[i], " "))
	}
}
