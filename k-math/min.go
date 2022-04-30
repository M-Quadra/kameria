package kmath

import (
	"math"
	"time"

	"golang.org/x/exp/constraints"
)

// Min default T
func Min[T constraints.Ordered](x ...T) T {
	var opt T
	if len(x) <= 0 {
		return opt
	}
	opt = x[0]

	for _, v := range x[1:] {
		switch any(opt).(type) {
		case float64:
			opt = any(math.Min(any(opt).(float64), any(v).(float64))).(T)
		default:
			if v < opt {
				opt = v
			}
		}
	}
	return opt
}

// MinAny default less: a<b
func MinAny[T any](less func(a, b T) bool, x ...T) T {
	var opt T
	if len(x) <= 0 {
		return opt
	}
	opt = x[0]

	for _, v := range x[1:] {
		if less(v, opt) {
			opt = v
		}
	}
	return opt
}

// MinTimes MinAny[time.Time]
func MinTimes(x ...time.Time) time.Time {
	return MinAny(func(a, b time.Time) bool {
		return a.Before(b)
	}, x...)
}
