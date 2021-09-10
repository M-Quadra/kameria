package kmath

type mean struct{}

// Mean support for {[]int|[]int64|[]float32|[]float64} -> float64
//  default 0
var Mean = mean{}

func (m mean) Ints(x []int) float64 {
	l := len(x)
	if l <= 0 {
		return 0
	}

	return Reduce(x, 0.0, func(result, elem interface{}) interface{} {
		return result.(float64) + float64(elem.(int))/float64(l)
	}).(float64)
}

func (m mean) Int64s(x []int64) float64 {
	l := len(x)
	if l <= 0 {
		return 0
	}

	return Reduce(x, 0.0, func(result, elem interface{}) interface{} {
		return result.(float64) + float64(elem.(int64))/float64(l)
	}).(float64)
}

func (m mean) Float32s(x []float32) float64 {
	l := len(x)
	if l <= 0 {
		return 0
	}

	return Reduce(x, 0.0, func(result, elem interface{}) interface{} {
		return result.(float64) + float64(elem.(float32))/float64(l)
	}).(float64)
}

func (m mean) Float64s(x []float64) float64 {
	l := len(x)
	if l <= 0 {
		return 0
	}

	return Reduce(x, 0.0, func(result, elem interface{}) interface{} {
		return result.(float64) + elem.(float64)/float64(l)
	}).(float64)
}
