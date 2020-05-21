package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func reverse(x int) int {
	isNegative := x < 0
	asString := strings.ReplaceAll(strconv.Itoa(x), "-", "")
	reversedString := ""

	for x := len(asString) - 1; x >= 0; x-- {
		reversedString += string(asString[x])
	}

	reversedInt, err := strconv.Atoi(reversedString)
	if err != nil || reversedInt < math.MinInt32 || reversedInt > math.MaxInt32 {
		// Unable to convert
		return 0
	}

	// If the original number was negative then re-apply it
	if isNegative {
		reversedInt *= -1
	}

	return reversedInt
}

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(120))
}
