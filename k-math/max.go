package kmath

import (
	"math"
	"time"

	"golang.org/x/exp/constraints"
)

// Max default T
func Max[T constraints.Ordered](x ...T) T {
	var opt T
	if len(x) <= 0 {
		return opt
	}
	opt = x[0]

	for _, v := range x[1:] {
		switch any(opt).(type) {
		case float64:
			opt = any(math.Max(any(opt).(float64), any(v).(float64))).(T)
		default:
			if v > opt {
				opt = v
			}
		}
	}
	return opt
}

// MaxAny default less: a<b
func MaxAny[T any](less func(a, b T) bool, x ...T) T {
	var opt T
	if len(x) <= 0 {
		return opt
	}
	opt = x[0]

	for _, v := range x[1:] {
		if less(opt, v) {
			opt = v
		}
	}
	return opt
}

// MaxTimes MaxAny[time.Time]
func MaxTimes(x ...time.Time) time.Time {
	return MaxAny(func(a, b time.Time) bool {
		return a.Before(b)
	}, x...)
}
