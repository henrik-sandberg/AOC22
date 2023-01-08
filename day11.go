package main

import (
	"fmt"
	"sort"
)

type monkey struct {
	items       []int
	operation   func(int) int
	target      func(int) int
	divisor     int
	inspections int
}

func Day11(input []string) {
	fmt.Println("Part 1: ", day11_part1())
	fmt.Println("Part 2: ", day11_part2())
}

func day11_part1() int {
	return solveDay11(getMonkeys(), 20, func(x int) int {
		return x / 3
	})
}

func day11_part2() int {
	monkeys := getMonkeys()
	divisor := 1
	for _, m := range monkeys {
		divisor *= m.divisor
	}
	return solveDay11(monkeys, 10000, func(x int) int {
		return x % divisor
	})
}

func solveDay11(monkeys map[int]*monkey, rounds int, decayFunc func(int) int) int {
	for round := 0; round < rounds; round++ {
		for i := 0; i < len(monkeys); i++ {
			m := monkeys[i]
			for _, item := range m.items {
				m.inspections += 1
				worryLevel := decayFunc(m.operation(item))
				next := monkeys[m.target(worryLevel%m.divisor)]
				next.items = append(next.items, worryLevel)
				m.items = m.items[1:]
			}
		}
	}
	insp := make([]int, len(monkeys))
	for _, monkey := range monkeys {
		insp = append(insp, monkey.inspections)
	}
	sort.Ints(insp)
	return insp[len(insp)-1] * insp[len(insp)-2]
}

func getMonkeys() map[int]*monkey {
	monkeys := map[int]*monkey{}
	monkeys[0] = &monkey{
		items:   []int{50, 70, 89, 75, 66, 66},
		divisor: 2,
		operation: func(x int) int {
			return x * 5
		},
		target: func(remainder int) int {
			if remainder == 0 {
				return 2
			} else {
				return 1
			}
		}}
	monkeys[1] = &monkey{
		items:   []int{85},
		divisor: 7,
		operation: func(x int) int {
			return x * x
		},
		target: func(remainder int) int {
			if remainder == 0 {
				return 3
			} else {
				return 6
			}
		}}
	monkeys[2] = &monkey{
		items:   []int{66, 51, 71, 76, 58, 55, 58, 60},
		divisor: 13,
		operation: func(x int) int {
			return x + 1
		},
		target: func(remainder int) int {
			if remainder == 0 {
				return 1
			} else {
				return 3
			}
		}}
	monkeys[3] = &monkey{
		items:   []int{79, 52, 55, 51},
		divisor: 3,
		operation: func(x int) int {
			return x + 6
		},
		target: func(remainder int) int {
			if remainder == 0 {
				return 6
			} else {
				return 4
			}
		}}
	monkeys[4] = &monkey{
		items:   []int{69, 92},
		divisor: 19,
		operation: func(x int) int {
			return x * 17
		},
		target: func(remainder int) int {
			if remainder == 0 {
				return 7
			} else {
				return 5
			}
		}}
	monkeys[5] = &monkey{
		items:   []int{71, 76, 73, 98, 67, 79, 99},
		divisor: 5,
		operation: func(x int) int {
			return x + 8
		},
		target: func(remainder int) int {
			if remainder == 0 {
				return 0
			} else {
				return 2
			}
		}}
	monkeys[6] = &monkey{
		items:   []int{82, 76, 69, 69, 57},
		divisor: 11,
		operation: func(x int) int {
			return x + 7
		},
		target: func(remainder int) int {
			if remainder == 0 {
				return 7
			} else {
				return 4
			}
		}}
	monkeys[7] = &monkey{
		items:   []int{65, 79, 86},
		divisor: 17,
		operation: func(x int) int {
			return x + 5
		},
		target: func(remainder int) int {
			if remainder == 0 {
				return 5
			} else {
				return 0
			}
		}}
	return monkeys
}
