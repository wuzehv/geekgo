1. 参考 Hystrix 实现一个滑动窗口计数器。

```
答：起初考虑用环形数组实现，但实现过程中发现无法重置计数，所以考虑用map来实现

实现结果：可通过传入参数来灵活配置窗口内容器个数；每秒一个容器

参考链接：https://pandaychen.github.io/2020/04/01/A-BUCKET-OF-HYSTRIX-ANALYSIS/
```