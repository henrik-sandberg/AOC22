package main

func Contains[T comparable](s []T, e T) bool {
    for _, v := range s {
        if v == e {
            return true
        }
    }
    return false
}

func Intersect[T comparable] (first []T, second []T) []T {
	result := []T{}
	for _, e := range first {
		if Contains(second, e) && !Contains(result, e) {
			result = append(result, e)
		}
	}
	return result
}

