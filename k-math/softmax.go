package kmath

import (
	"math"

	"golang.org/x/exp/constraints"
)

func softmax[T constraints.Float](logSoftmax bool, x ...T) []T {
	if len(x) <= 0 {
		return nil
	}
	max := Max(x...)

	sum := T(0)
	opt := make([]T, len(x))
	for i := range x {
		tmp := T(math.Exp(float64(x[i] - max)))
		if logSoftmax {
			opt[i] = x[i] - max
		} else {
			opt[i] = tmp
		}
		sum += tmp
	}
	if logSoftmax {
		sum = T(math.Log(float64(sum)))
	}

	for i := range opt {
		if logSoftmax {
			opt[i] = opt[i] - sum
		} else {
			opt[i] = opt[i] / sum
		}
	}
	return opt
}

// Softmax for Float
func Softmax[T constraints.Float](x ...T) []T {
	return softmax(false, x...)
}

// LogSoftmax for []Float
func LogSoftmax[T constraints.Float](x ...T) []T {
	return softmax(true, x...)
}
