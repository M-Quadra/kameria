package kameria

import (
	mathOld "math"

	kmath "github.com/M-Quadra/kameria/k-math"
)

// Math for call func
var Math = struct {
	Max kmath.Max
	Min kmath.Min
}{}

//Limit4Int return mid ∈ [min, max]
func Limit4Int(min, mid, max int) int {
	if min > max {
		min, max = max, min
	}

	mid = Math.Max.Ints(min, mid)
	mid = Math.Min.Ints(mid, max)
	return mid
}

//Limit4Int64 return mid ∈ [min, max]
func Limit4Int64(min, mid, max int64) int64 {
	if min > max {
		min, max = max, min
	}

	mid = Math.Max.Int64s(min, mid)
	mid = Math.Min.Int64s(mid, max)
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

// AvgFloat64 avg
func AvgFloat64(ary []float64) float64 {
	if len(ary) <= 0 {
		return 0
	}

	lenF := float64(len(ary))
	avg := 0.0
	for _, v := range ary {
		avg += v / float64(lenF)
	}
	return avg
}
