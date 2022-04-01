1. 基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。

```
答：errgroup的Go方法调用，其中有一个返回错误，内部会调用context的cancel方法

利用这个特性，在每个服务内监听context.Done，如果收到消息，说明可以进行退出
```