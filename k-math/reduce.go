package kmath

import "reflect"

// Reduce returns the result of combining the elements of the sequence using the given closure.
//   note init's Type
func Reduce(x interface{}, init interface{}, next func(result, elem interface{}) interface{}) interface{} {
	if x == nil {
		return nil
	}

	rv := reflect.ValueOf(x)
	if rv.Len() <= 0 {
		return nil
	}

	for i := 0; i < rv.Len(); i++ {
		v := rv.Index(i).Interface()
		init = next(init, v)
	}
	return init
}
