package algorithm

import "reflect"

// NextPermutation std::next_permutation
func NextPermutation(
	slice interface{}, first, last int,
	less func(i, j int) bool,
) bool {
	if slice = sliceCheck(slice); slice == nil || less == nil {
		return false
	}

	i := last - 1
	if first >= last || first == i {
		return false
	}

	for swap := reflect.Swapper(slice); ; {
		ip1 := i
		if i--; less(i, ip1) {
			j := last - 1
			for ; !less(i, j); j-- {
			}
			swap(i, j)
			Reverse(slice, ip1, last)
			return true
		}

		if i == first {
			Reverse(slice, first, last)
			return false
		}
	}
}
