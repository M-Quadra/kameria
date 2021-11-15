# algorithm

为了刷题时想用`STL`省事, 于是对部分功能进行拙劣的模仿

## 仿 std::reverse

```
algorithm.Reverse(slice interface{}, first, last int)
```

## 仿 std::next_permutation

```
algorithm.NextPermutation(
	slice interface{}, first, last int,
	less func(i, j int) bool,
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
