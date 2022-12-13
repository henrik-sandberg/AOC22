package main

import (
	"fmt"
)

func Day06(input []string) {
	fmt.Println("Part 1: ", day06_part1(input[0]))
	fmt.Println("Part 2: ", day06_part2(input[0]))
}

func day06_part1(input string) int {
	return findIndexOfDistinctCharacters(input, 4)
}

func day06_part2(input string) int {
	return findIndexOfDistinctCharacters(input, 14)
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
