# k-math

简易封装

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
kmath.Sum[T constraints.Integer | constraints.Float](x []T) T
```

使用统一的类型, 不考虑越界, 有特殊需求使用`Reduce`

# Reduce

```
kmath.Reduce[E any, R any](x []E, init R, next func(result R, elem E) R) R
```

序列遍历计算, 想法来自`Swift`高阶函数, 后面似乎得挪个地方

# 算术平均值

```
kmath.Mean[T constraints.Integer | constraints.Float](x []T) float64
```

使用`Reduce`实现, 默认0

# Softmax

```
kmath.Softmax[T constraints.Float](x ...T) []T
kmath.LogSoftmax[T constraints.Float](x ...T) []T
```

内部共享实现
