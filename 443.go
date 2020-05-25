package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Todo: finish this

func compress(chars *[]byte) int {

	charMap := make(map[byte]int, 0)
	for _, val := range *chars {
		if _, inMap := charMap[val]; inMap {
			charMap[val] += 1
		} else {
			charMap[val] = 1
		}
	}

	for k, v := range charMap {
		if v > 1 {
			*chars = []byte(strings.Replace(string(*chars), string(k), fmt.Sprintf("%s%s", string(k), strconv.Itoa(v)), 1))
		}
	}

	return len(*chars)
}

func main() {
	s := []byte("aabbccc")
	res := compress(&s)
	fmt.Println(string(s), res)
	//fmt.Println(compress([]byte("a")))
	//fmt.Println(compress([]byte("abbbbbbbbbbbb")))
}