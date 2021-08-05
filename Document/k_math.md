# k-math

# 最大值

```
kameria.Math.Max.Ints(x ...int) int
kameria.Math.Max.Int64s(x ...int64) int64
kameria.Math.Max.Strings(x ...string) string
kameria.Math.Max.Float32s(x ...float32) float32
kameria.Math.Max.Float64s(x ...float64) float64

kameria.Math.Max.Any(x interface{}, less func(a, b interface{}) bool) interface{}
```

直接使用小于号, 没有走`math.Max`

# 最小值

```
kameria.Math.Min.Ints(x ...int) int
kameria.Math.Min.Int64s(x ...int64) int64
kameria.Math.Min.Strings(x ...string) string
kameria.Math.Min.Float32s(x ...float32) float32
kameria.Math.Min.Float64s(x ...float64) float64

kameria.Math.Min.Any(x interface{}, less func(a, b interface{}) bool) interface{}
```

直接使用大于号, 没有走`math.Min`
