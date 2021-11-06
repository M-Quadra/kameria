# algorithm

为了处理特定情况下的蛋疼问题, 想用`STL`, 于是对部分函数进行拙劣的模范。

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
