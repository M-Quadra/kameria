# kameria

前途一片黑暗, 不知道未来还能不能养得起房东, 心好累

> 阿库娅大人说: "今天能养得起房东就今天养吧，反正明天也不一定养得起。就算养不起了，那也是世界的错！"
> 
> 我老婆说: "我讨厌的事有三件『办不到、好累、好麻烦』这三句话非常不好, 会抹杀人类所拥有的无限可能。"

随意堆砌, 企图发酵

# 胡闹

```
func StructTagMap(tag string, model interface{}) (map[string]interface{}, error)
```

映射`struct tag`对应元素的指针地址

后面应该会基于这个功能写一些方便操作

```
func IPv4StringToInt(ipStr string) (int64, error)
```

`IPv4`三点数表示法转10进制值地址


```
func IPv4IntToString(ipInt int64) (string, error)
```

10进制地址转`IPv4`三点数表示法

```
func Limit4Int(min, mid, max int) int
func Limit4Int64(min, mid, max int64) int64
func Limit4Float32(min, mid, max float32) float32
func Limit4Float64(min, mid, max float64) float64
```

范围取值

```
func uintptrAddress(v interface{}) (uintptr, error)
```

获取元素指针地址值, 暂不开放

```
Queue4String{}
```

string队列

```
func Href2Url(href, baseURL string) string
```

`href`转完整的`url`

```
func Float2String(i interface{}) string
```

浮点数转字符串, 格式为`'f'`

```
func SoftmaxFloat64(ary []float64) []float64
```

softmax

## []interface{}转化

```
kameria.SliceToInterfaces(slice interface{}) []interface{}
```

失败时返回nil

# Unique

## 切片去重

```
kameria.Unique.Ints(x []int) []int
kameria.Unique.Int64s(x []int64) []int64
kameria.Unique.Strings(x []string) []string
```

# 其他

[日常定时器](./Document/day_ticker.md)

[任务队列](./Document/operation_queue.md)

[k-math](./Document/k_math.md)

[algorithm](./Document/algorithm.md)

[[]byte 转 fs.FS](./virtual-fs/fs_test.go)