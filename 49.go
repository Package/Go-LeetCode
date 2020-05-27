package main

import (
	"fmt"
	"strings"
)

func isAnagram(w1 string, w2 string)  bool {
	if len(w1) != len(w2) { return false }

	for x := 0; x < len(w1); x++ {
		// Letter isn't in the string
		if strings.Index(w2, string(w1[x])) == -1 {
			return false
		}

		w2 = strings.Replace(w2, string(w1[x]), "", 1)
	}

	// If we get here then all the characters have passed so it's an anagram
	return true
}

func groupAnagrams(words []string) [][]string {
	anagrams := make([][] string, 0)

	for x := 0; x < len(words); x++ {
		anagram := false

		for y := 0; y < len(anagrams); y++ {
			if isAnagram(words[x], anagrams[y][0]) {
				anagram = true
				anagrams[y] = append(anagrams[y], words[x])
			}
		}

		// Not an anagram of any existing words so append this to the collection of anagrams as a new string[] entry
		if !anagram {
			anagrams = append(anagrams, []string{words[x]})
		}
	}

	return anagrams
}

func main() {
	fmt.Println(groupAnagrams([]string {"eat", "tea", "tan", "ate", "nat", "bat"}))
}
