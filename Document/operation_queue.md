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
添加任务, 当并行数不足时会阻塞, 并行数为`0`时同步执行

原本考虑加个队列缓存, 但是看到服务端队列累计后的资源占用, 还是算了
