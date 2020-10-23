package kameria

import (
	"math"
	"sort"
)

//Min4Int default 0
func Min4Int(elems ...int) int {
	if len(elems) <= 0 {
		return 0
	}

	sort.Ints(elems)
	return elems[0]
}

//Min4Int64 return min ∈ [min, max]
func Min4Int64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

//Min4Float32 return min ∈ [min, max]
func Min4Float32(a, b float32) float32 {
	return float32(math.Min(float64(a), float64(b)))
}

//Min4String dictionary order, default ""
func Min4String(elems ...string) string {
	if len(elems) <= 0 {
		return ""
	}

	sort.Strings(elems)
	return elems[0]
}

//Max4Int default 0
func Max4Int(elems ...int) int {
	if len(elems) <= 0 {
		return 0
	}

	sort.Ints(elems)
	return elems[len(elems)-1]
}

//Max4Int64 return max ∈ [min, max]
func Max4Int64(a, b int64) int64 {
	if a < b {
		return b
	}
	return a
}

//Max4Float32 return max ∈ [min, max]
func Max4Float32(a, b float32) float32 {
	return float32(math.Max(float64(a), float64(b)))
}

//Max4String dictionary order, default ""
func Max4String(elems ...string) string {
	if len(elems) <= 0 {
		return ""
	}

	sort.Strings(elems)
	return elems[len(elems)-1]
}

//Limit4Int return mid ∈ [min, max]
func Limit4Int(min, mid, max int) int {
	if min > max {
		min, max = max, min
	}

	mid = Max4Int(min, mid)
	mid = Min4Int(mid, max)
	return mid
}

//Limit4Int64 return mid ∈ [min, max]
func Limit4Int64(min, mid, max int64) int64 {
	if min > max {
		min, max = max, min
	}

	mid = Max4Int64(min, mid)
	mid = Min4Int64(mid, max)
	return mid
}

//Limit4Float32 return mid ∈ [min, max]
func Limit4Float32(min, mid, max float32) float32 {
	return float32(Limit4Float64(float64(min), float64(mid), float64(max)))
}

//Limit4Float64 return mid ∈ [min, max]
func Limit4Float64(min, mid, max float64) float64 {
	nMin := math.Min(min, max)
	nMax := math.Max(min, max)

	mid = math.Max(nMin, mid)
	mid = math.Min(mid, nMax)
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
		optAry = append(optAry, math.Exp(v))
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
