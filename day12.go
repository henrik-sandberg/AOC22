package main

import (
	"fmt"
)

func Day12(input []string) {
	root := buildGraphDay12(input)
	fmt.Println("Part 1: ", day12_part1(root))
	fmt.Println("Part 2: ", day12_part2())
}

func day12_part1(graph *node2) int {
	path := bfs(graph, "E")
	return len(path) - 1
}

func day12_part2() int {
	return 0
}

func buildGraphDay12(input []string) *node2 {
	fmt.Printf("Building graph from %dx%d input\n", len(input), len(input[0]))
	hills := map[point]*node2{}
	var root *node2
	for r_ind, row := range input {
		for c_ind, cell := range row {
			p := point{r_ind, c_ind}
			hills[p] = &node2{value: calculateElevation(cell)}
			// fmt.Println("Building graph at point", p, "with", hills[p])
			if cell == 83 {
				hills[p].name = "S"
				root = hills[p]
				fmt.Println("Found starting point", root, "at", p)
			} else if cell == 69 {
				hills[p].name = "E"
				fmt.Println("Found target", hills[p], "at", p)
			} else {
				hills[p].name = fmt.Sprint(p)
			}
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
			//fmt.Println("Begin comparing for cell", cell, "found at point", p, found)
			for _, dir := range dirs {
				if other, found := hills[point{r_ind + dir.rd, c_ind + dir.cd}]; found {
					appendValidNode(currentNode, other)
				}
			}
		}
	}
	fmt.Println("Returning graph root", root)
	return root
}

// appends b to edges of a, if possible to go from a to b
func appendValidNode(a *node2, b *node2) {
	if a.value >= b.value-1 {
		a.edges = append(a.edges, b)
	}
}

func calculateElevation(r rune) int {
	if r == 83 {
		return 0
	} else if r == 69 {
		return int('z' - 'a')
	} else {
		return int(r - 'a')
	}
}
