package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
)

func Day13(input []string) {
	fmt.Println("Part 1: ", day13_part1(input))
	fmt.Println("Part 2: ", day13_part2(input))
}

func day13_part1(input []string) (result int) {
	for i := 0; i < len(input); i = i + 3 {
		var left []any
		var right []any
		json.Unmarshal([]byte(input[i]), &left)
		json.Unmarshal([]byte(input[i+1]), &right)
		if day13_compare(left, right) < 0 {
			index := i/3 + 1
			result += index
		}
	}
	return
}

func day13_part2(input []string) int {
	packets := [][]any{}
	for _, line := range input {
		if line != "" {
			var packet []any
			json.Unmarshal([]byte(line), &packet)
			packets = append(packets, packet)
		}
	}
	dividerPackets := [][]any{{[]any{float64(2)}}, {[]any{float64(6)}}}
	packets = append(packets, dividerPackets...)

	sort.Slice(packets, func(i, j int) bool {
		return day13_compare(packets[i], packets[j]) < 0
	})

	// Ugly string conversion as any does not implement comparable
	strings := []string{}
	for _, p := range packets {
		strings = append(strings, fmt.Sprint(p))
	}

	indexes := []int{}
	for _, dp := range dividerPackets {
		indexes = append(indexes, IndexOf(strings, fmt.Sprint(dp))+1)
	}

	return multiply(indexes...)
}

func day13_compare(left []any, right []any) int {
	for i := 0; i < len(left) && i < len(right); i++ {
		l := left[i]
		r := right[i]
		typeLeft := reflect.TypeOf(l).Kind()
		if typeLeft == reflect.TypeOf(r).Kind() {
			if typeLeft == reflect.Float64 {
				if res := compareFloat64s(l.(float64), r.(float64)); res != 0 {
					return res
				}
			} else if res := day13_compare(l.([]any), r.([]any)); res != 0 {
				return res
			}
		} else {
			var res int
			if typeLeft == reflect.Float64 {
				res = day13_compare([]any{l.(float64)}, r.([]any))
			} else {
				res = day13_compare(l.([]any), []any{r.(float64)})
			}
			if res != 0 {
				return res
			}
		}
	}
	return compareInts(len(left), len(right))
}
