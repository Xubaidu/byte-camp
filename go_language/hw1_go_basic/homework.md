# homework

## desc

1. 修改第一个例子猜谜游戏里面的最终代码，使用 `fmt.Scanf` 来简化代码实现
2. 修改第二个例子命令行词典里面的最终代码，增加另一种翻译引擎的支持
3. 在上一步骤的基础上，修改代码实现并行请求两个翻译引擎来提高响应速度

## sol

1. 格式化输入即可，参考 [guess.go](/guess-game/guess.go)
2. 加入了火山翻译，参考 [dict.go](/simple-dict/dict.go)
3. 用 `sync.WaitGroup` 和 `go func` 实现