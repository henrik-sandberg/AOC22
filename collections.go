package main

type void struct{}

// Get all unique values from the slice as a map
func Set[T comparable](s []T) map[T]void {
	result := make(map[T]void)
	var v void
	for _, val := range s {
		result[val] = v
	}
	return result
}

func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

func Intersect[T comparable](first []T, second []T) []T {
	result := []T{}
	for _, e := range first {
		if Contains(second, e) && !Contains(result, e) {
			result = append(result, e)
		}
	}
	return result
}
