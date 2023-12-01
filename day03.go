package main

import (
	"fmt"
)

func Day03(input []string) {
	fmt.Println("Part 1: ", day03_part1(input))
	fmt.Println("Part 2: ", day03_part2(input))
}

func day03_part1(input []string) (result int) {
	for _, bag := range input {
		runeArray := []rune(bag)
		intersect := Intersect(runeArray[:len(runeArray)/2], runeArray[len(runeArray)/2:])
		result += calculatePriority(intersect)
	}
	return
}

func day03_part2(input []string) (result int) {
	for i := 0; i < len(input); i += 3 {
		common := Intersect([]rune(input[i]), []rune(input[i+1]))
		common = Intersect(common, []rune(input[i+2]))
		result += calculatePriority(common)
	}
	return
}

func calculatePriority(chars []rune) int {
	if len(chars) != 1 {
		fmt.Printf("Invalid number of chars in %v\n", chars)
	}
	num := int(chars[0])
	if num >= 97 {
		num -= 96
	} else {
		num -= 38
	}
	return num
}
