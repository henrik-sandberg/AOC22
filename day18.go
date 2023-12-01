package main

import (
	"fmt"
	"strconv"
	"strings"
)

type cube struct {
	x, y, z int
}

func Day18(input []string) {
	cubes := map[cube]bool{}
	for _, line := range input {
		c := strings.Split(line, ",")
		x, _ := strconv.Atoi(c[0])
		y, _ := strconv.Atoi(c[1])
		z, _ := strconv.Atoi(c[2])
		cubes[cube{x, y, z}] = true
	}
	fmt.Println("Result part 1:", day18_part1(cubes))
	fmt.Println("Result part 2:", day18_part2(cubes))
}

func day18_part1(cubes map[cube]bool) (res int) {
	for c := range cubes {
		for side := range sides(c) {
			if !cubes[side] {
				res += 1
			}
		}
	}
	return
}

func day18_part2(cubes map[cube]bool) (res int) {
	var c cube
	queue := []cube{{-1, -1, -1}}
	seen := map[cube]bool{}
	for len(queue) > 0 {
		c, queue = queue[0], queue[1:]
		for side := range sides(c) {
			if cubes[side] {
				res += 1
			} else if !seen[side] &&
				c.x >= -1 && c.x < 21 &&
				c.y >= -1 && c.y < 21 &&
				c.z >= -1 && c.z < 21 {
				seen[side] = true
				queue = append(queue, side)
			}
		}
	}
	return
}

func sides(c cube) <-chan cube {
	ch := make(chan cube)
	go func() {
		ch <- cube{c.x + 1, c.y, c.z}
		ch <- cube{c.x - 1, c.y, c.z}
		ch <- cube{c.x, c.y + 1, c.z}
		ch <- cube{c.x, c.y - 1, c.z}
		ch <- cube{c.x, c.y, c.z + 1}
		ch <- cube{c.x, c.y, c.z - 1}
		close(ch)
	}()
	return ch
}
