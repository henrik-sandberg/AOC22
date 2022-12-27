package main

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func max(n ...int) int {
	high := n[0]
	for _, v := range n[1:] {
		if v > high {
			high = v
		}
	}
	return high
}
