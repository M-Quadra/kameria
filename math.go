package kameria

import (
	mathOld "math"
	"reflect"
	"time"

	kmath "github.com/M-Quadra/kameria/k-math"
)

type math struct{}

// Math for call kmath func simply without type
var Math = math{}

// only support for ...{int|int64|string|float32|float64|time.Time}
//  default nil
func (m math) Max(x ...interface{}) interface{} {
	if len(x) <= 0 {
		return nil
	}

	switch x[0].(type) {
	case int:
		return kmath.Max.Any(x, func(a, b interface{}) bool {
			return a.(int) < b.(int)
		})
	case int64:
		return kmath.Max.Any(x, func(a, b interface{}) bool {
			return a.(int64) < b.(int64)
		})
	case string:
		return kmath.Max.Any(x, func(a, b interface{}) bool {
			return a.(string) < b.(string)
		})
	case float32:
		return kmath.Max.Any(x, func(a, b interface{}) bool {
			return a.(float32) < b.(float32)
		})
	case float64:
		return kmath.Max.Any(x, func(a, b interface{}) bool {
			return a.(float64) < b.(float64)
		})
	case time.Time:
		return kmath.Max.Any(x, func(a, b interface{}) bool {
			return a.(time.Time).Before(b.(time.Time))
		})
	}

	return nil
}

// only support for ...{int|int64|string|float32|float64|time.Time}
//  default nil
func (m math) Min(x ...interface{}) interface{} {
	if len(x) <= 0 {
		return nil
	}

	switch x[0].(type) {
	case int:
		return kmath.Min.Any(x, func(a, b interface{}) bool {
			return a.(int) > b.(int)
		})
	case int64:
		return kmath.Min.Any(x, func(a, b interface{}) bool {
			return a.(int64) > b.(int64)
		})
	case string:
		return kmath.Min.Any(x, func(a, b interface{}) bool {
			return a.(string) > b.(string)
		})
	case float32:
		return kmath.Min.Any(x, func(a, b interface{}) bool {
			return a.(float32) > b.(float32)
		})
	case float64:
		return kmath.Min.Any(x, func(a, b interface{}) bool {
			return a.(float64) > b.(float64)
		})
	case time.Time:
		return kmath.Min.Any(x, func(a, b interface{}) bool {
			return a.(time.Time).After(b.(time.Time))
		})
	}

	return nil
}

// only support for []{int|int64|float32|float64}
//  default nil
func (m math) Sum(x interface{}) interface{} {
	rv := reflect.ValueOf(x)
	if rv.Len() <= 0 {
		return nil
	}

	switch rv.Index(0).Interface().(type) {
	case int:
		return kmath.Reduce(x, int(0), func(result, elem interface{}) interface{} {
			return result.(int) + elem.(int)
		}).(int)
	case int64:
		return kmath.Reduce(x, int64(0), func(result, elem interface{}) interface{} {
			return result.(int64) + elem.(int64)
		})
	case float32:
		return kmath.Reduce(x, float32(0), func(result, elem interface{}) interface{} {
			return result.(float32) + elem.(float32)
		})
	case float64:
		return kmath.Reduce(x, float64(0), func(result, elem interface{}) interface{} {
			return result.(float64) + elem.(float64)
		})
	}

	return nil
}

// only support for []{int|int64|float32|float64}
//  default 0
func (m math) Mean(x interface{}) float64 {
	rv := reflect.ValueOf(x)
	l := rv.Len()
	if l <= 0 {
		return 0
	}

	switch rv.Index(0).Interface().(type) {
	case int:
		return kmath.Reduce(x, float64(0), func(result, elem interface{}) interface{} {
			return result.(float64) + float64(elem.(int))/float64(l)
		}).(float64)
	case int64:
		return kmath.Reduce(x, float64(0), func(result, elem interface{}) interface{} {
			return result.(float64) + float64(elem.(int64))/float64(l)
		}).(float64)
	case float32:
		return kmath.Reduce(x, float64(0), func(result, elem interface{}) interface{} {
			return result.(float64) + float64(elem.(float32))/float64(l)
		}).(float64)
	case float64:
		return kmath.Reduce(x, float64(0), func(result, elem interface{}) interface{} {
			return result.(float64) + elem.(float64)/float64(l)
		}).(float64)
	}

	return 0
}

//Limit4Int return mid ∈ [min, max]
func Limit4Int(min, mid, max int) int {
	if min > max {
		min, max = max, min
	}

	mid = kmath.Max.Ints(min, mid)
	mid = kmath.Min.Ints(mid, max)
	return mid
}

//Limit4Int64 return mid ∈ [min, max]
func Limit4Int64(min, mid, max int64) int64 {
	if min > max {
		min, max = max, min
	}

	mid = kmath.Max.Int64s(min, mid)
	mid = kmath.Min.Int64s(mid, max)
	return mid
}

//Limit4Float32 return mid ∈ [min, max]
func Limit4Float32(min, mid, max float32) float32 {
	return float32(Limit4Float64(float64(min), float64(mid), float64(max)))
}

//Limit4Float64 return mid ∈ [min, max]
func Limit4Float64(min, mid, max float64) float64 {
	nMin := mathOld.Min(min, max)
	nMax := mathOld.Max(min, max)

	mid = mathOld.Max(nMin, mid)
	mid = mathOld.Min(mid, nMax)
	return mid
}

// SoftmaxFloat64 softmax
func SoftmaxFloat64(ary []float64) []float64 {
	if ary == nil {
		return nil
	}
	if len(ary) <= 0 {
		return []float64{}
	}

	optAry := []float64{}

	sum := 0.0
	for i, v := range ary {
		optAry = append(optAry, mathOld.Exp(v))
		sum += optAry[i]
	}
	for i := range optAry {
		optAry[i] /= sum
	}
	return optAry
}
