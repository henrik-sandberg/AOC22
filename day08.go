package main

import (
	"fmt"
)

func Day08(input []string) {
	forest := parseForest(input)

	fmt.Println("Part 1:", day08_part1(forest))
	fmt.Println("Part 2:", day08_part2(forest))
}

func day08_part1(forest [][]int) int {
	visible := make([][]int, len(forest))
	for ri, row := range forest {
		visible[ri] = make([]int, len(row))
		visible[ri][0] = 1
		max := row[0]
		for ci := 1; ci < len(row); ci++ {
			if row[ci] > max {
				visible[ri][ci] |= 1
				max = row[ci]
			}
		}
		max = row[len(row)-1]
		visible[ri][len(row)-1] = 1
		for ci := len(row) - 1; ci >= 0; ci-- {
			if row[ci] > max {
				visible[ri][ci] |= 1
				max = row[ci]
			}
		}
	}
	for ci := 0; ci < len(forest[0]); ci++ {
		visible[0][ci] = 1
		max := forest[0][ci]
		for ri := 1; ri < len(forest)-1; ri++ {
			if forest[ri][ci] > max {
				visible[ri][ci] |= 1
				max = forest[ri][ci]
			}
		}
		visible[len(forest)-1][ci] = 1
		max = forest[len(forest)-1][ci]
		for ri := len(forest) - 1; ri >= 0; ri-- {
			if forest[ri][ci] > max {
				visible[ri][ci] |= 1
				max = forest[ri][ci]
			}
		}
	}
	return sum(visible)
}

func day08_part2(forest [][]int) int {
	max := 0
	for r := range forest {
		for c := range forest[r] {
			if score := calculateScenicScore(r, c, forest); score > max {
				max = score
			}
		}
	}
	return max
}

func calculateScenicScore(row int, col int, forest [][]int) int {
	h := len(forest)
	w := len(forest[row])
	val := forest[row][col]
	res := 1

	up := 0
	for i := row - 1; i >= 0; i-- {
		up++
		if val <= forest[i][col] {
			break
		}
	}
	res *= up

	down := 0
	for i := row + 1; i < h; i++ {
		down++
		if val <= forest[i][col] {
			break
		}
	}
	res *= down

	left := 0
	for i := col - 1; i >= 0; i-- {
		left++
		if val <= forest[row][i] {
			break
		}
	}
	res *= left

	right := 0
	for i := col + 1; i < w; i++ {
		right++
		if val <= forest[row][i] {
			break
		}
	}
	res *= right
	return res
}

func parseForest(input []string) [][]int {
	ret := make([][]int, len(input))
	for ri, row := range input {
		ret[ri] = make([]int, len(row))
		for ci, c := range row {
			ret[ri][ci] = int(c - '0')
		}
	}
	return ret
}
