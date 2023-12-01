package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Day04(input []string) {
	score_part1 := 0
	score_part2 := 0

	for _, line := range input {
		first, second := toIntPairs(line)
		score_part1 += day04_part1(first, second)
		score_part2 += day04_part2(first, second)
	}
	fmt.Println("Part 1: ", score_part1)
	fmt.Println("Part 2: ", score_part2)
}

func day04_part1(first []int, second []int) int {
	if first[0] >= second[0] && first[1] <= second[1] || second[0] >= first[0] && second[1] <= first[1] {
		return 1
	}
	return 0
}

func day04_part2(first []int, second []int) int {
	if first[0] >= second[0] && first[0] <= second[1] ||
		first[1] >= second[0] && first[1] <= second[1] ||
		second[0] >= first[0] && second[0] <= first[1] ||
		second[1] >= first[0] && second[1] <= first[1] {
		return 1
	}
	return 0
}

func toIntPairs(s string) ([]int, []int) {
	pair := strings.Split(s, ",")

	first := strings.Split(pair[0], "-")
	first_low, _ := strconv.Atoi(first[0])
	first_high, _ := strconv.Atoi(first[1])

	second := strings.Split(pair[1], "-")
	second_low, _ := strconv.Atoi(second[0])
	second_high, _ := strconv.Atoi(second[1])

	return []int{first_low, first_high}, []int{second_low, second_high}
}
