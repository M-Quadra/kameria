package kameria

import "math"

//Max4Int return max ∈ [min, max]
func Max4Int(a, b int) int {
	if a < b {
		return b
	}
	return a
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

//Min4Int return min ∈ [min, max]
func Min4Int(a, b int) int {
	if a < b {
		return a
	}
	return b
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
