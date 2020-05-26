package main

import (
	"fmt"
)

type point struct {
	X int
	Y int
}

func findStartingPositions(board [][]byte, firstChar string) []point {
	startingPoints := make([]point, 0)

	for x := 0; x < len(board); x++ {
		for y := 0; y < len(board[x]); y++ {
			if string(board[x][y]) == firstChar {
				startingPoints = append(startingPoints, point{x, y})
			}
		}
	}

	return startingPoints
}

func exist(board [][]byte, word string) bool {
	starting := findStartingPositions(board, string(word[0]))

	for _, sp := range starting {
		if canMakeWordStartingAt(sp.X, sp.Y, board, word, 0) {
			return true
		}
	}

	return false
}

// Performs a recursive call to check whether the words exist
// This is an example of a DFS (depth first search) algorithm.
func canMakeWordStartingAt(x int, y int, board [][]byte, s string, index int) bool {
	if index == len(s) {
		// Made the whole word
		return true
	}

	// Out of bounds check
	if x < 0 || x > len(board) - 1 || y < 0 || y > len(board[x]) - 1 {
		//fmt.Println("Outta bounds")
		return false
	}

	// Found a path that words, make S shorter
	if string(board[x][y]) == string(s[index]) {
		tmp := board[x][y]
		board[x][y] = ' '

		// Check in the other directions
		next := canMakeWordStartingAt(x + 1, y, board, s, index + 1) ||
				canMakeWordStartingAt(x - 1, y, board, s, index + 1) ||
				canMakeWordStartingAt(x, y - 1, board, s, index + 1) ||
				canMakeWordStartingAt(x, y + 1, board, s, index + 1)

		board[x][y] = tmp

		return next
	}

	return false
}

func main() {
	s := [][]byte {{'A','B','C','E'},{'S','F','C','S'},{'A','D','E','E'}}
	fmt.Println(exist(s, "ABCCED"))
	fmt.Println(exist(s, "SEE"))
	fmt.Println(exist(s, "ABCB"))
}
