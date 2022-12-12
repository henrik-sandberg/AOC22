package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type node struct {
	name string
	children map[string]*node
	value int
}


func Day07(input []string) {
	root := buildGraph(input)
	resultPart1, usedSpace := day07_part1(&root)
	fmt.Println("Part 1: ", resultPart1)

	unusedSpace := 70000000 - usedSpace
	spaceToFree := 30000000 - unusedSpace
	resultPart2, _ := day07_part2(&root, spaceToFree)
	fmt.Println("Part 2: ", resultPart2)
}

// Calculate result of part 1
// Returns a tuple of result (if satisfying the requirement of size limit)
// and the running total of space used in subfolders
func day07_part1(n *node) (int, int) {
	limit := 100000
	result := 0
	runningTotal := n.value
	for _, v := range n.children {
		res, total := day07_part1(v)
		result += res
		runningTotal += total
	}
	if runningTotal <= limit {
		result += runningTotal
	}
	return result, runningTotal
}

// Calculate result of part 2
// Returns a tuple of best result per folder (or MaxInt if not satisfying the requirement)
// and the running total of space used in subfolders
func day07_part2(n *node, spaceToFree int) (int, int) {
	result := math.MaxInt
	runningTotal := n.value
	for _, v := range n.children {
		res, total := day07_part2(v, spaceToFree)
		runningTotal += total
		if res >= spaceToFree && res < result {
			result = res
		}
	}
	if runningTotal >= spaceToFree && runningTotal < result {
		result = runningTotal
	}
	return result, runningTotal
}

func buildGraph(input[] string) node {
	stack := []*node{}
	for _, line := range input {
		arr := strings.Split(line, " ")
		if arr[0] == "$" {
			if cmd := arr[1]; cmd == "cd" {
				if arr[2] == ".." {
					stack = stack[:len(stack)-1]
				} else if arr[2] == "/" {
					root := node{name: "/", children: map[string]*node{}}
					stack = append(stack, &root)
				} else {
					directory := stack[len(stack)-1].children[arr[2]]
					stack = append(stack, directory)
				}
			}
		} else if arr[0] == "dir" {
			name := arr[1]
			newNode := node{name: name, children: map[string]*node{}}
			stack[len(stack)-1].children[name] = &newNode 
		} else {
			size, _ := strconv.Atoi(arr[0])
			stack[len(stack)-1].value += size
		}
	}
	return *stack[0]
}

