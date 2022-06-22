package kmath

// Reduce returns the result of combining the elements of the sequence using the given closure.
func Reduce[E any, R any](x []E, init R, next func(result R, elem E) R) R {
	if len(x) <= 0 {
		return init
	}

	for _, v := range x {
		init = next(init, v)
	}
	return init
}
