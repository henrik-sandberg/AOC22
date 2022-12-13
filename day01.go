package main

import (
	"fmt"
	"sort"
	"strconv"
)

func Day01(input []string) {
	counters := make([]int, 1)
	for _, s := range input {
		if s == "" {
			counters = append(counters, 0)
		} else {
			i, _ := strconv.Atoi(s)
			counters[len(counters)-1] += i
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(counters)))
	fmt.Printf("Top 1 calories: %d\n", counters[0])
	fmt.Printf("Top 3 calories: %d\n", counters[0]+counters[1]+counters[2])
}
