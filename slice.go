package kameria

import (
	"reflect"
)

// SliceToInterfaces slice -> []interface{}
func SliceToInterfaces(slice interface{}) []interface{} {
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

	opt := make([]interface{}, 0, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		opt = append(opt, rv.Index(i).Interface())
	}
	return opt
}
