package kameria

import (
	"reflect"
)

//Unique4String string ary
func Unique4String(iptAry []string) []string {
	uniAry := unique(iptAry)
	optAry := make([]string, len(uniAry))

	for i, v := range uniAry {
		optAry[i] = v.(string)
	}
	return optAry
}

//Unique4Int int ary
func Unique4Int(iptAry []int) []int {
	uniAry := unique(iptAry)
	optAry := make([]int, len(uniAry))

	for i, v := range uniAry {
		optAry[i] = v.(int)
	}
	return optAry
}

//Unique4int64 int64 ary
func Unique4int64(iptAry []int64) []int64 {
	uniAry := unique(iptAry)
	optAry := make([]int64, len(uniAry))

	for i, v := range uniAry {
		optAry[i] = v.(int64)
	}
	return optAry
}

func unique(ary interface{}) []interface{} {
	rv := reflect.ValueOf(ary)
	tmpMap := map[interface{}]struct{}{}
	optAry := []interface{}{}

	for i := 0; i < rv.Len(); i++ {
		v := rv.Index(i).Interface()
		_, ok := tmpMap[v]
		if ok {
			continue
		}

		tmpMap[v] = struct{}{}
		optAry = append(optAry, v)
	}

	return optAry
}
