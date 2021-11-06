package algorithm

import (
	"reflect"
)

func sliceCheck(slice interface{}) interface{} {
	if slice == nil {
		return nil
	}

	rv := reflect.ValueOf(slice)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Slice {
		return nil
	}

	return rv.Interface()
}

// Reverse std::reverse
func Reverse(slice interface{}, first, last int) {
	if slice = sliceCheck(slice); slice == nil {
		return
	}
	swap := reflect.Swapper(slice)

	for last--; first < last; first, last = first+1, last-1 {
		swap(first, last)
	}
}
