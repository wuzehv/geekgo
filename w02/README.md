1. dao层的sql.ErrNoRows是否应该Wrap这个error，抛给上层，为什么？应该怎么做？

```
答：应该wrap，这个问题的核心是dao层，如果dao层不对错误进行wrap，那么调用方就会对底层实现产生依赖
```