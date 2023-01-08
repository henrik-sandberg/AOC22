package main

// Returns all unique values from the slice as a map (with nil values)
func Set[T comparable](s []T) (result map[T]interface{}) {
	for _, val := range s {
		result[val] = nil
	}
	return
}

// Returns true if s contains e
func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

// Returns the intersect of first and second as a new slice
func Intersect[T comparable](first []T, second []T) (result []T) {
	for _, e := range first {
		if Contains(second, e) && !Contains(result, e) {
			result = append(result, e)
		}
	}
	return
}
