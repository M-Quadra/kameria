# kameria

闭门造毂, 乃不知有轮

```
func StructTagMap(tag string, model interface{}) (map[string]interface{}, error)
```

映射`struct tag`对应元素的指针地址

后面应该会基于这个功能写一些方便操作

```
func HasError(err error) bool
```

错误检查与打印, 追溯3层

```
func IPv4StringToInt(ipStr string) (int64, error)
```

`IPv4`三点数表示法转10进制值地址


```
func IPv4IntToString(ipInt int64) (string, error)
```

10进制地址转`IPv4`三点数表示法

```
func Max4Int(a, b int) int
func Max4Int64(a, b int64) int64
func Max4Float32(a, b float32) float32
```

取较大值

```
func Min4Int(a, b int) int
func Min4Int64(a, b int64) int64
func Min4Float32(a, b float32) float32
```

取较小值

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
func Unique4String(iptAry []string) []string
func Unique4Int(iptAry []int) []int
func Unique4int64(iptAry []int64) []int64
```

数组去重

```
Queue4String{}
```

string队列