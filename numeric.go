package main

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func min(n ...int) (min int) {
	min = n[0]
	for _, v := range n[1:] {
		if v < min {
			min = v
		}
	}
	return
}

func max(n ...int) (max int) {
	max = n[0]
	for _, v := range n[1:] {
		if v > max {
			max = v
		}
	}
	return
}

func sum(arr [][]int) int {
	result := 0
	for _, row := range arr {
		for _, cell := range row {
			result += cell
		}
	}
	return result
}
