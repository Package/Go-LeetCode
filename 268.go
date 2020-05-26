package main

import "fmt"

func missingNumber(nums []int) int {
	seen := make([]int, len(nums) + 1)
	for _, value := range nums {
		seen[value] = 1
	}

	for index, value := range seen {
		if value == 0 {
			return index
		}
	}

	return -1
}

func main() {
	fmt.Println(missingNumber([]int {3,0,1}))
	fmt.Println(missingNumber([]int {9,6,4,2,3,5,7,0,1}))
}
