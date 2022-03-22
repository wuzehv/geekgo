1. dao层的sql.ErrNoRows是否应该Wrap这个error，抛给上层，为什么？应该怎么做？

```
答：应该wrap，提供更多调试信息，例如原生sql
```