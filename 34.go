package main

import "fmt"

func searchRange(data []int, target int) []int {
	firstPos := -1
	lastPost := -1

	for x := 0; x < len(data); x++ {
		// We can safely assume once we have seen a number bigger than the target
		// that the target will never re-appear as the data is provided in a sorted format.
		if data[x] > target { break }

		if data[x] == target {
			if firstPos == -1 {
				// First time we have encountered the target
				firstPos = x
				lastPost = x
			} else {
				// Any subsequent time is the last time
				lastPost = x
			}
		}
	}

	// If either first or last was not found, the problem wants us to return both as -1 to indicate this
	if firstPos == -1 || lastPost == -1 {
		firstPos = -1
		lastPost = -1
	}
	return []int{firstPos, lastPost}
}

func main() {
	fmt.Println(searchRange([]int{1,2,3,3,3,3,4,5,9}, 3))
}
