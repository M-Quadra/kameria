# k-math

原本期望在内外均能保持相同的调用路径, 如

```
kameria.Math.Max.Ints(x ...int) int
kmath.Max.Ints(x ...int) int
```

并且保持不暴露非必要方法或结构体。

一是思来想去发觉没法实现, 二来`kameria.Math.Max.Type`的调用链路过长

索性决定外部`kameria`使用`interface{}`, 内部`kmath`使用具体类型。（所以啥时候有泛型啊？

# 最大值

```
kameria.Math.Max(x ...interface{}) interface{}

kmath.Max.Ints(x ...int) int
kmath.Max.Int64s(x ...int64) int64
kmath.Max.Strings(x ...string) string
kmath.Max.Float32s(x ...float32) float32
kmath.Max.Float64s(x ...float64) float64
kmath.Max.Times(x ...time.Time) time.Time

kmath.Max.Any(x interface{}, less func(a, b interface{}) bool) interface{}
```

判断方式为`<`, 没有走`math.Max`

# 最小值

```
kameria.Math.Min(x ...interface{}) interface{}

kmath.Min.Ints(x ...int) int
kmath.Min.Int64s(x ...int64) int64
kmath.Min.Strings(x ...string) string
kmath.Min.Float32s(x ...float32) float32
kmath.Min.Float64s(x ...float64) float64
kmath.Min.Times(x ...time.Time) time.Time

kmath.Min.Any(x interface{}, large func(a, b interface{}) bool) interface{}
```

判断方式为`>`, 没有走`math.Min`

# 求和

```
kameria.Math.Sum(x interface{}) interface{}

kmath.Sum.Ints(x []int) int
kmath.Sum.Int64s(x []int64) int64
kmath.Sum.Float32s(x []float32) float32
kmath.Sum.Float64s(x []float64) float64
```

使用统一的类型, 不考虑越界, 有特殊需求使用`Reduce`

# Reduce

```
kmath.Reduce(x interface{}, init interface{}, next func(result, elem interface{}) interface{}) interface{}
```

序列遍历计算, 想法来自`Swift`高阶函数, 后面似乎得挪个地方
