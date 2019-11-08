package kameria

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

//Limit4Int return mid ∈ [min, max]
func Limit4Int(min, mid, max int) int {
	if min > max {
		min ^= max
		max ^= min
		min ^= max
	}

	mid = Max4Int(min, mid)
	mid = Min4Int(mid, max)
	return mid
}

//Limit4Int64 return mid ∈ [min, max]
func Limit4Int64(min, mid, max int64) int64 {
	if min > max {
		min ^= max
		max ^= min
		min ^= max
	}

	mid = Max4Int64(min, mid)
	mid = Max4Int64(mid, max)
	return mid
}
