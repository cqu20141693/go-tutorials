### cobra

[使用指南](https://github.com/spf13/cobra/blob/master/user_guide.md)
https://darjun.github.io/2020/01/17/godailylib/cobra/
#### run

```
Cobra 提供了两种方法来运行我们的逻辑：Run func(cmd *Command, args []string) 和 RunE func(cmd *Command, args []string) error ，
后者可以返回一个错误，我们将能够从 Execute() 方法的返回中捕获。
```