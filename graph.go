package main

import (
	"fmt"
)

type point struct {
	x, y int
}

type node struct {
	name  string
	edges []*node
	value int
}

func (n *node) getEdge(name string) (*node, bool) {
	for _, edge := range n.edges {
		if edge.name == name {
			return edge, true
		}
	}
	return &node{}, false
}

func (n *node) String() string {
	edges := make([]string, len(n.edges))
	for _, edge := range n.edges {
		edges = append(edges, edge.name)
	}
	return fmt.Sprintf("Node{ name: %s, edges: %d (%v) }", n.name, len(n.edges), edges)
}

// Finds the shortest path of a graph from specified starting point. Ends when reaches a node named "target".
// If unsolvable, shortestPath is an empty slice.
func bfs(start *node, target string) (shortestPath []string) {
	queue := [][]*node{{start}}
	visited := map[string]bool{start.name: true}
	for len(queue) > 0 {
		path := queue[0]
		n := path[len(path)-1]
		queue = queue[1:]
		for _, edge := range n.edges {
			if edge.name == target {
				for _, edge := range path {
					shortestPath = append(shortestPath, edge.name)
				}
				shortestPath = append(shortestPath, edge.name)
				return
			}
			if !visited[edge.name] {
				visited[edge.name] = true
				newPath := []*node{}
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
