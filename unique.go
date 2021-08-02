package kameria

import (
	"reflect"
)

type unique int

// Unique support []string, []int, []int64
const Unique = unique(0)

func (u unique) unique(x interface{}) map[interface{}]struct{} {
	rv := reflect.ValueOf(x)

	opt := make(map[interface{}]struct{}, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		opt[rv.Index(i).Interface()] = struct{}{}
	}
	return opt
}

func (u unique) Ints(x []int) []int {
	if x == nil {
		return nil
	}

	m := u.unique(x)
	opt := make([]int, 0, len(m))
	for _, v := range x {
		if _, ok := m[v]; ok {
			opt = append(opt, v)
			delete(m, v)
		}
	}

	return opt
}

func (u unique) Int64s(x []int64) []int64 {
	if x == nil {
		return nil
	}

	m := u.unique(x)
	opt := make([]int64, 0, len(m))
	for _, v := range x {
		if _, ok := m[v]; ok {
			opt = append(opt, v)
			delete(m, v)
		}
	}

	return opt
}

func (u unique) Strings(x []string) []string {
	if x == nil {
		return nil
	}

	m := u.unique(x)
	opt := make([]string, 0, len(m))
	for _, v := range x {
		if _, ok := m[v]; ok {
			opt = append(opt, v)
			delete(m, v)
		}
	}

	return opt
}
