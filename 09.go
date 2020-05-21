package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	if x < 0 { return false }

	asString := strconv.Itoa(x)
	halfway := (len(asString) / 2) + 1

	for i := 0; i < halfway; i++ {
		x1 := string(asString[i])
		x2 := string(asString[len(asString)-1-i])
		if x1 != x2 { return false }
	}

	return true
}

func main() {
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(9001009))
}