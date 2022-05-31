package algorithm

import (
	"reflect"
)

// Reverse std::reverse
func Reverse[T any](slice []T, first, last int) {
	if len(slice) <= 0 {
		return
	}

	swap := reflect.Swapper(slice)
	for last--; first < last; first, last = first+1, last-1 {
		swap(first, last)
	}
}
