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
添加任务, 当并行数不足时会阻塞

原本考虑加个队列缓存, 但是看到服务端队列累计后的资源占用, 还是算了

```
.Close(async ...bool)
```

因为任务阻塞是依靠`chan`实现的, 这会导致`goroutine`数量逐步上升

对于`Job`或全局控制来说其实无所谓, 但为了实现更细致的控制就加上关闭方法

可以选择是否异步关闭, 但由于内存问题, 在`que:=OperationQueue{}`的创建方式下, 如果作死调用`que.Close(true);que=OperationQueue{};que.Close()`开辟新队列再关闭会引起崩溃

使用`que=&OperationQueue{}`可以解决这个问题, 但我不想多写代码, 目前也没什么优雅的解决办法, 反正最大并行数可以修改没必要刻意作死