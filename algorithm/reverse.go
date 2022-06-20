package algorithm

// Reverse std::reverse
func Reverse[T any](slice []T, first, last int) {
	if len(slice) <= 0 {
		return
	}

	for last--; first < last; first, last = first+1, last-1 {
		slice[first], slice[last] = slice[last], slice[first]
	}
}
