# k-math

原本期望在内外均能保持相同的调用路径, 如

```
kameria.Math.Max.Ints(x ...int) int
kmath.Max.Ints(x ...int) int
```

并且保持不暴露非必要方法或结构体。

一是思来想去发觉没法实现, 二来`kameria.Math.Max.Type`的调用链路过长

泛型不是很完善, `switch`后依然过不了编译, 保留几个特定类型的方法

# 最大值

```
kmath.Max[T constraints.Ordered](x ...T) T
kmath.MaxAny[T any](less func(a, b T) bool, x ...T) T
kmath.MaxTimes(x ...time.Time) time.Time
```

`less`默认`<`, `float64`走`math.Max`

# 最小值

```
kmath.Min[T constraints.Ordered](x ...T) T
kmath.MinAny[T any](less func(a, b T) bool, x ...T) T
kmath.MinTimes(x ...time.Time) time.Time
```

`less`默认`<`, `float64`走`math.Min`

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

# 算术平均值

```
kameria.Math.Mean(x interface{}) float64

kmath.Mean.Ints(x []int) float64
kmath.Mean.Int64s(x []int64) float64
kmath.Mean.Float32s(x []float32) float64
kmath.Mean.Float64s(x []float64) float64
```

使用`Reduce`实现, 默认0