package main

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func min(n ...int) int {
	min := n[0]
	for _, v := range n[1:] {
		if v < min {
			min = v
		}
	}
	return min
}

func max(n ...int) int {
	max := n[0]
	for _, v := range n[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

func multiply(n ...int) int {
	product := n[0]
	for _, v := range n[1:] {
		product *= v
	}
	return product
}

func sum(arr [][]int) (result int) {
	for _, row := range arr {
		for _, cell := range row {
			result += cell
		}
	}
	return
}

// Compares two ints. Returns negative number if a is less than b, positive number if a is bigger than b; zero otherwise
func compareInts(a int, b int) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}

// Compares two float64s. Returns negative number if a is less than b, positive number if a is bigger than b; zero otherwise
func compareFloat64s(a float64, b float64) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}
