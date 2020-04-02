# OperationQueue

任务队列, 疑似生产者/消费者模型, 照着`NSOperationQueue`命名的

```
.MaxConcurrentOperationCount(cnt ...int) int
```

最大并行数, 由于Go无法整合Get/Set, 用可选参数曲线救国

无值时Get, 有值时Set

```
.AddOperation(fc func())
```

添加任务, 当并行数不足时会进入等待队列