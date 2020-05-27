package main

import "fmt"

var (
	FreshOrange  = 1
	RottenOrange = 2
)

type orangeFlip struct {
	X int
	Y int
}

func orangesRotting(grid [][]int) int {
	iterations := 0

	for {
		flips := getOrangeFlips(grid)
		if len(flips) == 0 {
			break
		}

		for _, value := range flips {
			grid[value.X][value.Y] = RottenOrange
		}

		iterations++
	}


	// Check there are no fresh oranges left, if there are then this input cannot be solved so return -1
	if countFreshOranges(grid) > 0 {
		return -1
	}

	return iterations
}

// Counts how many fresh oranges are in the grid.
// We use this at the end of the problem to determine whether it was actually solvable or not.
// If there is a count > 0 then the problem was not solvable.
func countFreshOranges(grid [][]int) int {
	oranges := 0

	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			if grid[x][y] == FreshOrange {
				oranges += 1
			}
		}
	}

	return oranges
}

// Iterates the grid returning any points at which a cell should be turned rotten. (adjacent to a previously rotten cell)
func getOrangeFlips(i [][]int) []orangeFlip {

	// Tracks the positions of which oranges should be flipped during each cycle
	flips := make([]orangeFlip, 0)

	for x := 0; x < len(i); x++ {
		for y := 0; y < len(i[x]); y++ {
			curr := i[x][y]
			if curr != RottenOrange { continue }

			if x > 0 {
				if i[x-1][y] == FreshOrange { // Left
					flips = append(flips, orangeFlip{x-1, y})
				}
			}
			if x < len(i) - 1 {
				if i[x+1][y] == FreshOrange { // Right
					flips = append(flips, orangeFlip{x+1, y})
				}
			}
			if y > 0 {
				if i[x][y-1] == FreshOrange { // Up
					flips = append(flips, orangeFlip{x, y-1})
				}
			}
			if y < len(i[x]) - 1 {
				if i[x][y+1] == FreshOrange { // Down
					flips = append(flips, orangeFlip{x, y+1})
				}
			}
		}
	}

	return flips
}

func main() {
	fmt.Println(orangesRotting([][]int { {2, 1, 1}, {1, 1, 0}, {0, 1, 1}}))
	fmt.Println(orangesRotting([][]int { {2, 1, 1}, {0, 1, 1}, {1, 0, 1}}))
	fmt.Println(orangesRotting([][]int { {0, 2} }))
}
