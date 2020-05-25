package main

import (
	"fmt"
	"sort"
)

func sortMapIntoString(characterMap *map[string]int) string {
	type keyValue struct {
		Key string
		Value int
	}

	charSlice := make([]keyValue, 0)
	for k, v := range *characterMap {
		charSlice = append(charSlice, keyValue{k, v})
	}

	sort.Slice(charSlice, func(x, y int) bool {
		return charSlice[x].Value > charSlice[y].Value
	})

	sortedString := ""
	for _, kv := range charSlice {
		for x := 0; x < kv.Value; x++ {
			sortedString += kv.Key
		}
	}

	return sortedString
}

func buildFrequencyMap(s string) *map[string]int {
	characterMap := map[string]int {}

	for _, char := range s {
		asString := string(char)

		if _, inMap := characterMap[asString]; inMap {
			// If it's already in the map then increase the frequency
			characterMap[asString] += 1
		} else {
			// Not in the map yet so add it
			characterMap[asString] = 1
		}
	}

	return &characterMap
}

func frequencySort(s string) string {
	if len(s) == 0 {
		return s
	}

	charMap := buildFrequencyMap(s)
	sortedString := sortMapIntoString(charMap)

	return sortedString
}

func main() {
	fmt.Println(frequencySort("tree"))
	fmt.Println(frequencySort("cccaaa"))
	fmt.Println(frequencySort("Aabb"))
	fmt.Println(frequencySort("A"))
}
