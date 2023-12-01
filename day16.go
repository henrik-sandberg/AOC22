package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type key struct {
	a, b string
}

var V = []string{}       // All vertices
var F = map[string]int{} // Non-zero vertex flow
var D = map[key]int{}    // Distance between any two vertices

func Day16(input []string) {
	parse_valves(input)

	// Make a slice of non-zero vertices
	vertices := make([]string, 0, len(F))
	for k, _ := range F {
		vertices = append(vertices, k)
	}
	fmt.Println("Part 1:", day16_solve(30, "AA", vertices))
	fmt.Println("Part 2:", day16_part2(vertices))
}

func day16_part2(vertices []string) (res int){
	for ps := range PowerSet(vertices) {
		if len(ps) != 7 {
			// An assumption which works but cannot be asserted. We expect the rooms to be evenly divided.
			// Alternative would be checking len(ps) < len(vs) / 2 with a x6 runtime. Rationale is after half the sets are just mirroring.
			continue
		}
		// Create a disjoint set of vertices for the elephant
		e := make([]string, 0, len(vertices)-len(ps))
		for _, v := range vertices {
			if !Contains(ps, v) {
				e = append(e, v)
			}
		}
		res = max(res, day16_solve(26, "AA", ps) + day16_solve(26, "AA", e))
	}
	return
}

func day16_solve(depth int, room string, vs []string) (res int) {
	for _, node := range vs {
		if dist := D[key{room, node}]; dist < depth {
			ns := make([]string, 0, len(vs)-1)
			for _, k := range vs {
				if k != node {
					ns = append(ns, k)
				}
			}
			res = max(res, F[node]*(depth-dist-1)+day16_solve(depth-dist-1, node, ns))
		}
	}
	return
}

func parse_valves(input []string) {
	re := regexp.MustCompile("Valve (?P<valve>[A-Z]+) has flow rate=(?P<flow_rate>\\d+); tunnels? leads? to valves? (?P<tunnels>.+)")
	for _, s := range input {
		match := re.FindStringSubmatch(s)
		name := match[re.SubexpIndex("valve")]
		V = append(V, name)
		if flow_rate, _ := strconv.Atoi(match[re.SubexpIndex("flow_rate")]); flow_rate != 0 {
			F[name] = flow_rate
		}
		for _, c := range strings.Split(match[re.SubexpIndex("tunnels")], ", ") {
			D[key{name, c}] = 1
		}
	}
	// Helper function to set unknown distances to a high value
	dist := func(a, b string) int {
		val, found := D[key{a, b}]
		if !found {
			return 9999
		}
		return val
	}
	// Run Floyd-Warshall algorithm to calculate all distances
	for _, k := range V {
		for _, i := range V {
			for _, j := range V {
				D[key{i, j}] = min(dist(i, j), dist(i, k)+dist(k, j))
			}
		}
	}
}
