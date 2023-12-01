package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Day07(input []string) {
	root := buildGraph(input)
	fmt.Println("Part 1: ", day07_part1(&root))
	fmt.Println("Part 2: ", day07_part2(&root, root.value-40000000))
}

func day07_part1(n *node) int {
	result := 0
	if n.value <= 100000 {
		result = n.value
	}
	for _, v := range n.nodes {
		result += day07_part1(v)
	}
	return result
}

func day07_part2(n *node, spaceToFree int) int {
	result := math.MaxInt
	if n.value < result && n.value >= spaceToFree {
		result = n.value
	}
	for _, v := range n.nodes {
		if res := day07_part2(v, spaceToFree); res < result && res >= spaceToFree {
			result = res
		}
	}
	return result
}

func buildGraph(input []string) node {
	stack := []*node{}
	for _, line := range input {
		arr := strings.Split(line, " ")
		if arr[0] == "$" {
			if cmd := arr[1]; cmd == "cd" {
				if target := arr[2]; target == ".." {
					stack = stack[:len(stack)-1]
				} else if target == "/" {
					root := node{name: target}
					stack = append(stack, &root)
				} else {
					directory, _ := stack[len(stack)-1].getNode(target)
					stack = append(stack, directory)
				}
			}
		} else if arr[0] == "dir" {
			name := arr[1]
			e := &stack[len(stack)-1].nodes
			(*e) = append((*e), &node{name: name})
		} else {
			size, _ := strconv.Atoi(arr[0])
			// Back propagating file size to root makes filtering much easier later
			for _, n := range stack {
				n.value += size
			}
		}
	}
	return *stack[0]
}
