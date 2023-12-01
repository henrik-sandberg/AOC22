package main

import (
	"fmt"
)

type point struct {
	x, y int
}

type node struct {
	name  string
	nodes []*node
	value int
}

// Get the linked node. Returns true or false as secondary variable if found or not.
func (n *node) getNode(name string) (*node, bool) {
	for _, n := range n.nodes {
		if n.name == name {
			return n, true
		}
	}
	return &node{}, false
}

func (n *node) String() string {
	nodes := make([]string, len(n.nodes))
	for _, n := range n.nodes {
		nodes = append(nodes, n.name)
	}
	return fmt.Sprintf("Node{ name: %s, edges: %d (%v) }", n.name, len(n.nodes), n)
}

// Finds the shortest path of a graph from specified starting point.
// Ends when reaches a node named "target".
// Returns a slice of the node names found being the shortest path
// If unsolvable, shortestPath is an empty slice.
func bfs(start *node, target string) (shortestPath []string) {
	queue := [][]*node{{start}}
	visited := map[string]bool{start.name: true}
	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		for _, n := range path[len(path)-1].nodes {
			if n.name == target {
				for _, n := range path {
					shortestPath = append(shortestPath, n.name)
				}
				shortestPath = append(shortestPath, n.name)
				return
			}
			if !visited[n.name] {
				visited[n.name] = true
				newPath := []*node{}
				for _, n := range path {
					newPath = append(newPath, n)
				}
				newPath = append(newPath, n)
				queue = append(queue, newPath)
			}
		}
	}
	return
}
