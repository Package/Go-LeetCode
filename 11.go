package main

import (
	"fmt"
	"math"
)

func bestWaterFromPos(pos int, elements []int) int {
	bestOutput := 0

	for x := pos; x < len(elements); x++ {
		a := int(math.Min(float64(elements[pos]), float64(elements[x])))
		b := x - pos

		if a * b > bestOutput {
			bestOutput = a * b
		}
	}

	return bestOutput
}


func maxArea(height []int) int {

	bestOutcome := 0
	for x := 0; x < len(height); x++ {
		currOutcome := bestWaterFromPos(x, height)

		if currOutcome > bestOutcome {
			bestOutcome = currOutcome
		}
	}

	return bestOutcome
}

func main() {
	fmt.Println(maxArea([]int {1,8,6,2,5,4,8,3,7}))
}