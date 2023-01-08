package main

import (
	"fmt"
)

type point struct {
	x, y int
}
type node struct {
	name  string
	edges map[string]*node
	value int
}

// slimmer implementation with edges being in a slice
type node2 struct {
	name  string
	edges []*node2
	value int
}

func (n *node2) String() string {
	edges := make([]string, len(n.edges))
	for _, edge := range n.edges {
		edges = append(edges, edge.name)
	}
	return fmt.Sprintf("Node{ name: %s, edges: %d (%v) }", n.name, len(n.edges), edges)
}

func bfs(start *node2, target string) (shortestPath []string) {
	queue := [][]*node2{{start}}
	visited := map[string]bool{start.name: true}
	for len(queue) > 0 {
		path := queue[0]
		curNode := path[len(path)-1]
		queue = queue[1:]
		for _, edge := range curNode.edges {
			if edge.name == target {
				for _, edge := range path {
					shortestPath = append(shortestPath, edge.name)
				}
				shortestPath = append(shortestPath, edge.name)
				return
			}
			if !visited[edge.name] {
				visited[edge.name] = true
				newPath := []*node2{}
				for _, n := range path {
					newPath = append(newPath, n)
				}
				newPath = append(newPath, edge)
				queue = append(queue, newPath)
			}
		}
	}
	return
}
