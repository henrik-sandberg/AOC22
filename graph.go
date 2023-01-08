package main

import (
	"fmt"
	"strings"
)

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

func bfs(start *node2, target string) (path []string) {
	fmt.Println("Solving with BFS for graph start", start.name, "and target", target)
	queue := [][]*node2{{start}}
	visited := map[string]bool{start.name: true}
	for len(queue) > 0 {
		currentPath := queue[0]
		currentNode := currentPath[len(currentPath)-1]
		queue = queue[1:]

		if strings.HasPrefix(currentNode.name, "{0 ") {
			keys := make([]string, len(visited))
			for key := range visited {
				keys = append(keys, key)
			}
			fmt.Println("\n*******\nAT TOP ROW", currentNode.name)
		}
		fmt.Printf("Currently at: %s\nQueue length: %d\nVisited nodes: %d\nConnected edges: %v\nComing from path:%v\nAll paths:\n", currentNode.name, len(queue), len(visited), nodeNames(currentNode.edges), nodeNames(currentPath))
		for _, q := range queue {
			fmt.Println(nodeNames(q))
		}

		fmt.Println()
		for _, edge := range currentNode.edges {
			if edge.name == target {
				for _, edge := range currentPath {
					path = append(path, edge.name)
				}
				path = append(path, edge.name)
				fmt.Println("Found target!", path)
				return
			}
			if !visited[edge.name] {
				visited[edge.name] = true
				newPath := append(currentPath, edge)
				queue = append(queue, newPath)
			}
		}
	}
	fmt.Println("No path found")
	return
}

func nodeNames(nodes []*node2) (res []string) {
	for _, n := range nodes {
		res = append(res, n.name)
	}
	return
}
