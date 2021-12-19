package algorithm

// BinaryBound std::lower_bound or std::upper_bound
//  chose (middle, last) when comp return true
//  chose [first, middle) when comp return false
//  lower_bound comp example, { return a[middle] < v }
//  upper_bound comp example, { return a[middle] <= v }
func BinaryBound(first, last int, comp func(middle int) bool) int {
	for count, half, middle := last-first, 0, 0; count > 0; {
		half = count >> 1
		middle = first + half
		if comp(middle) {
			first = middle + 1
			count -= half + 1
			continue
		}

		count = half
	}
	return first
}
