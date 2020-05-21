package main

import "fmt"

func twoSum(nums[] int, target int) []int {
	for xIndex, xVal := range nums {
		for yIndex, yVal := range nums {
			if xVal + yVal == target && xIndex != yIndex {
				return []int {xIndex, yIndex}
			}
		}
	}

	panic("Solution not possible")
}

func main() {
	fmt.Println(twoSum([]int {2, 7, 11, 15}, 9))
}
