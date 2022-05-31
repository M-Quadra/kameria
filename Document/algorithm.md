# algorithm

为了刷题时想用`STL`省事, 于是对部分功能进行拙劣的模仿。

为了调用直观, 似乎应该把package直接换成`std`?

## 仿 std::reverse

```
algorithm.Reverse[T any](slice []T, first, last int)
```

## 仿 std::next_permutation

```
algorithm.NextPermutation[T any](
	slice []T, first, last int,
	less func(a, b T) bool,
) bool
```

## 仿 std::priority_queue

```
priorityqueue.New(slicePtr interface{}, less func(i, j int) bool) (*PriorityQueue, error)
```

初始化, 需要外部切片, 使用`container/heap`实现,有朝一日若有范型考虑随缘改造

```
pq.Size() int
pq.Empty() bool
pq.Push(x interface{})
pq.Pop() interface{}
pq.Top() interface{}
```

`.Pop()`多了返回值, 其余没啥变化

## 仿 std::lower_bound or std::upper_bound

```
algorithm.BinaryBound(first, last int, comp func(middle int) bool) int
```

由于go并没有运算符重载, 想要开放比较方法。如此一来, `std::lower_bound`与`std::upper_bound`就相互显得有些重复了, 故将二者合并。
