package algorithm

// NextPermutation std::next_permutation
func NextPermutation[T any](
	slice []T, first, last int,
	less func(a, b T) bool,
) bool {
	if len(slice) <= 0 || less == nil {
		return false
	}

	i := last - 1
	if first >= last || first == i {
		return false
	}

	for {
		ip1 := i
		if i--; less(slice[i], slice[ip1]) {
			j := last - 1
			for ; !less(slice[i], slice[j]); j-- {
			}
			slice[i], slice[j] = slice[j], slice[i]
			Reverse(slice, ip1, last)
			return true
		}

		if i == first {
			Reverse(slice, first, last)
			return false
		}
	}
}
