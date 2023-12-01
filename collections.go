package main

// Returns all unique values from the slice as a map (with nil values)
func Set[T comparable](s []T) map[T]interface{} {
	result := make(map[T]interface{})
	for _, val := range s {
		result[val] = nil
	}
	return result
}

// Returns the index of e in s. -1 if s does not contain e
func IndexOf[T comparable](s []T, e T) int {
	for i, v := range s {
		if v == e {
			return i
		}
	}
	return -1
}

// Returns true if s contains e
func Contains[T comparable](s []T, e T) bool {
	return IndexOf(s, e) != -1
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
