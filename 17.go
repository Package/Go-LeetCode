package main

import (
	"fmt"
)

// Todo: this is incomplete


var PhoneMap = map[string]string {
	"2": "abc", "3": "def", "4": "ghi", "5": "jkl", "6": "mno",
	"7": "pqrs", "8": "tuv", "9": "wxyz",
}

func letterCombinations(digits string) []string {
	combos := []string {}

	for index, v := range digits {
		asString := string(v)
		fmt.Printf("%d %s\n", index, asString)
	}

	return combos
}

func main() {
	fmt.Println(letterCombinations("23"))
}
