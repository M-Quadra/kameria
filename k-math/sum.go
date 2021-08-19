package kmath

type sum int

// Sum for []{int}
const Sum = sum(0)

func (s sum) Ints(x []int) int {
	if len(x) <= 0 {
		return 0
	}

	v := Reduce(x, 0, func(result, elem interface{}) interface{} {
		return result.(int) + elem.(int)
	})
	return v.(int)
}

func (s sum) Int64s(x []int64) int64 {
	if len(x) <= 0 {
		return 0
	}

	v := Reduce(x, int64(0), func(result, elem interface{}) interface{} {
		return result.(int64) + elem.(int64)
	})
	return v.(int64)
}

func (s sum) Float32s(x []float32) float32 {
	if len(x) <= 0 {
		return 0
	}

	v := Reduce(x, float32(0), func(result, elem interface{}) interface{} {
		return result.(float32) + elem.(float32)
	})
	return v.(float32)
}

func (s sum) Float64s(x []float64) float64 {
	if len(x) <= 0 {
		return 0
	}

	v := Reduce(x, float64(0), func(result, elem interface{}) interface{} {
		return result.(float64) + elem.(float64)
	})
	return v.(float64)
}
