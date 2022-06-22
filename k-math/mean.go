package kmath

import "golang.org/x/exp/constraints"

// Mean return float64, default 0
func Mean[T constraints.Integer | constraints.Float](x []T) float64 {
	l := len(x)
	if l <= 0 {
		return 0
	}

	return Reduce(x, 0.0, func(result float64, elem T) float64 {
		return result + float64(elem)/float64(l)
	})
}
