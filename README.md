# kameria

闭门造毂, 乃不知有轮, 胡乱弄了些代码

前途一片黑暗, 不知道未来还能不能养得起房东, 心好累

> 阿库娅大人说: "今天能养得起房东就今天养吧，反正明天也不一定养得起。就算养不起了，那也是世界的错！"
> 
> 我老婆说: "我讨厌的事有三件『办不到、好累、好麻烦』这三句话非常不好, 会抹杀人类所拥有的无限可能。"

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
func Min4Int(elems ...int) int
func Min4Int64(a, b int64) int64
func Min4Float32(a, b float32) float32
```

取较小值

```
func Min4String(elems ...string) string
```

取字典序最小值

```
func Max4Int(elems ...int) int
func Max4Int64(a, b int64) int64
func Max4Float32(a, b float32) float32
```

取较大值

```
func Max4String(elems ...string) string
```

取字典序最大值

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

```
func JSONStrToMap(jsonStr string) (map[string]interface{}, error)
```

json 转 map[string]interface{}

```
func Int2String(i interface{}, base int) string
```

int 转 string, 内置了类型判断

```
func Arabic2NumStr(arabic string) string
```

阿拉伯-印度文数字`۰۱۲۳۴۵۶۷۸۹`转`0123456789`

```
func Tibetan2NumStr(arabic string) string
```

藏文数字`༠༡༢༣༤༥༦༧༨༩`转`0123456789`

```
func Href2Url(href, baseURL string) string
```

`href`转完整的`url`

# 其他

[通用设置](.Document/config.md)

[日常定时器](.Document/day_ticker.md)

[任务队列](./Document/operation_queue.md)
