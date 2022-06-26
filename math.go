package kameria

import (
	mathOld "math"

	kmath "github.com/M-Quadra/kameria/k-math"
)

type math struct{}

// Math for call kmath func simply without type
var Math = math{}

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
