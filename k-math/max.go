package kmath

import (
	"reflect"
)

// Max for ...{int|int64|string|float32|float64}
type Max int

// Any x must be Array, Chan, Map, Slice, or String
func (m Max) Any(x interface{}, less func(a, b interface{}) bool) interface{} {
	if x == nil {
		return nil
	}

	rv := reflect.ValueOf(x)
	if rv.Len() <= 0 {
		return nil
	}

	maxN := rv.Index(0).Interface()
	for i := 1; i < rv.Len(); i++ {
		n := rv.Index(i).Interface()
		if less(maxN, n) {
			maxN = n
		}
	}
	return maxN
}

// Ints default 0
func (m Max) Ints(x ...int) int {
	if len(x) <= 0 {
		return 0
	}

	v := m.Any(x, func(a, b interface{}) bool {
		return a.(int) < b.(int)
	})
	return v.(int)
}

// Int64s default 0
func (m Max) Int64s(x ...int64) int64 {
	if len(x) <= 0 {
		return 0
	}

	v := m.Any(x, func(a, b interface{}) bool {
		return a.(int64) < b.(int64)
	})
	return v.(int64)
}

// Strings default ""
func (m Max) Strings(x ...string) string {
	if len(x) <= 0 {
		return ""
	}

	v := m.Any(x, func(a, b interface{}) bool {
		return a.(string) < b.(string)
	})
	return v.(string)
}

// Float32s default 0
func (m Max) Float32s(x ...float32) float32 {
	if len(x) <= 0 {
		return 0
	}

	v := m.Any(x, func(a, b interface{}) bool {
		return a.(float32) < b.(float32)
	})
	return v.(float32)
}

// Float64s default 0
func (m Max) Float64s(x ...float64) float64 {
	if len(x) <= 0 {
		return 0
	}

	v := m.Any(x, func(a, b interface{}) bool {
		return a.(float64) < b.(float64)
	})
	return v.(float64)
}
