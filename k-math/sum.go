package kmath

import "golang.org/x/exp/constraints"

// Sum for []{Integer | Float}
func Sum[T constraints.Integer | constraints.Float](x []T) T {
	var zero T
	return Reduce(x, zero, func(result, elem T) T {
		return result + elem
	})
}
