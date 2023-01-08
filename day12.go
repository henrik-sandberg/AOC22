package main

import (
	"fmt"
	"math"
)

func Day12(input []string) {
	graph := buildGraphDay12(input)
	fmt.Println("Part 1: ", day12_part1(graph))
	fmt.Println("Part 2: ", day12_part2(graph))
}

func day12_part1(graph []*node2) int {
	for _, n := range graph {
		if n.name == "S" {
			path := bfs(n, "E")
			return len(path) - 1
		}
	}
	fmt.Println("No starting point found")
	return -1
}

func day12_part2(graph []*node2) int {
	res := math.MaxInt32
	for _, n := range graph {
		if n.value == 0 {
			if path := bfs(n, "E"); len(path) > 0 {
				res = min(res, len(path)-1)
			}
		}
	}
	return res
}

func buildGraphDay12(input []string) (nodes []*node2) {
	fmt.Printf("Building graph from %dx%d input\n", len(input), len(input[0]))
	hills := map[point]*node2{}
	for r_ind, row := range input {
		for c_ind, cell := range row {
			p := point{r_ind, c_ind}
			hills[p] = buildNode(cell, p)
		}
	}
	dirs := []struct{ rd, cd int }{
		{rd: 0, cd: 1},
		{rd: 0, cd: -1},
		{rd: 1, cd: 0},
		{rd: -1, cd: 0},
	}
	for r_ind, row := range input {
		for c_ind := range row {
			currentNode := hills[point{r_ind, c_ind}]
			for _, dir := range dirs {
				if other, found := hills[point{r_ind + dir.rd, c_ind + dir.cd}]; found {
					appendValidNode(currentNode, other)
				}
			}
		}
	}
	for _, v := range hills {
		nodes = append(nodes, v)
	}
	return
}

// appends b to edges of a, if possible to go from a to b
func appendValidNode(a *node2, b *node2) {
	if a.value >= b.value-1 {
		a.edges = append(a.edges, b)
	}
}

func buildNode(r rune, p point) *node2 {
	var name string
	var elevation int
	if r == 83 {
		name = "S"
		elevation = 0
	} else if r == 69 {
		name = "E"
		elevation = int('z' - 'a')
	} else {
		name = fmt.Sprint(p)
		elevation = int(r - 'a')
	}
	return &node2{value: elevation, name: name}
}
