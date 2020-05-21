package main

import "fmt"


func sequenceStartingAt(s string, startAt int) int {
	seen := map[string]string {}
	sequence := 0

	for x:=startAt; x < len(s); x++ {
		asString := string(s[x])
		_, alreadySeenCharacter := seen[asString]
		if alreadySeenCharacter {
			// We've met a repeating character so return
			break
		}

		sequence += 1
		seen[asString] = asString
	}

	return sequence
}

func lengthOfLongestSubstring(s string) int {

	bestSeq := 0

	for i, _ := range s {
		curr := sequenceStartingAt(s, i)
		if curr > bestSeq {
			bestSeq = curr
		}
	}

	return bestSeq
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
