package main

import (
	"fmt"
)

var BestLength int = 0

func isUniqueString(s string) bool {
	charArr := make([]int, 26)

	for _, value := range s {
		if charArr[value - 'a'] > 0 {
			return false
		}

		charArr[value - 'a'] = 1
	}

	return true
}

func findMaxLength(arr []string, index int, currStr string) {
	if len(currStr) > BestLength && isUniqueString(currStr) {
		BestLength = len(currStr)
	}

	if index == len(arr) {
		return
	}

	findMaxLength(arr, index + 1, currStr)
	findMaxLength(arr, index + 1, currStr + arr[index])
}

func maxLength(arr []string) int {
	findMaxLength(arr, 0, "")
	tmpBest := BestLength
	BestLength = 0
	return tmpBest
}

func main() {
	fmt.Println(maxLength([]string {"un","iq","ue"}))
	fmt.Println(maxLength([]string {"cha","r","act","ers"}))
	fmt.Println(maxLength([]string {"abcdefghijklmnopqrstuvwxyz"}))
	fmt.Println(maxLength([]string {"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p"}))
}
