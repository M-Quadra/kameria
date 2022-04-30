package kameria

import (
	mathOld "math"
	"reflect"

	kmath "github.com/M-Quadra/kameria/k-math"
)

type math struct{}

// Math for call kmath func simply without type
var Math = math{}

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

	mid = kmath.Max(min, mid)
	mid = kmath.Min(mid, max)
	return mid
}

//Limit4Int64 return mid ∈ [min, max]
func Limit4Int64(min, mid, max int64) int64 {
	if min > max {
		min, max = max, min
	}

	mid = kmath.Max(min, mid)
	mid = kmath.Min(mid, max)
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
