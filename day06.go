package main

import (
	"fmt"
)

func Day06(input []string) {
	fmt.Println("Part 1: ", findIndexOfDistinctCharacters(input[0], 4))
	fmt.Println("Part 2: ", findIndexOfDistinctCharacters(input[0], 14))
}

// Find index where the subsequent characters are all distinct
func findIndexOfDistinctCharacters(s string, n int) int {
	for i := n; i < len(s); i++ {
		if set := Set([]rune(s[i-n : i])); len(set) == n {
			return i
		}
	}
	return 0
}
