# DayTicker

一个简易定时器, 阻塞至一天内的特定时间, 若当天时间无法满足要求则推迟到下一天

```
NewDayTicker(hr, min, sec int) *DayTicker
```

创建一个定时器

```
.Hour(hr ...int) int
.Minute(min ...int) int
.Second(sec ...int) int
```

时间设置, 利用可变参数实现`Get/Set`

```
.Next() time.Time
```

阻塞至下一个目标时间, 为防止短时间大量调用, 最短阻塞时间`1s`