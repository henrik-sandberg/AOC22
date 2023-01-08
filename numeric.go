package main

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
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
