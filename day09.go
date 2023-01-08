package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Day09(input []string) {
	fmt.Println("Part 1:", solveDay09(input, 2))
	fmt.Println("Part 2:", solveDay09(input, 10))
}

func solveDay09(input []string, length int) int {
	seen := []point{}
	rope := make([]point, length)
	for _, cmd := range input {
		cmdArr := strings.Split(cmd, " ")
		direction := cmdArr[0]
		distance, _ := strconv.Atoi(cmdArr[1])
		for n := 0; n < distance; n++ {
			rope[0].move(direction)
			for i := 1; i < len(rope); i++ {
				rope[i].follow(rope[i-1])
			}
			seen = append(seen, rope[len(rope)-1])
		}
	}
	return len(Set(seen))
}

func (p *point) move(direction string) {
	switch direction {
	case "U":
		p.y += 1
	case "D":
		p.y -= 1
	case "R":
		p.x += 1
	case "L":
		p.x -= 1
	}
}

func (p *point) follow(head point) {
	if abs(head.x-p.x) > 1 || abs(head.y-p.y) > 1 {
		if head.x > p.x {
			p.x += 1
		} else if head.x < p.x {
			p.x -= 1
		}
		if head.y > p.y {
			p.y += 1
		} else if head.y < p.y {
			p.y -= 1
		}
	}
}
