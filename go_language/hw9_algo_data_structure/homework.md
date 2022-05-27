# 实现 v1 版本的 pdqsort

## 实现

实现 v1 版本的 pdqsort，见 [main.go](main.go)

## 测试

测试代码参考 [mysort/all_test.go](mysort/all_test.go)

测试命令如下

```shell
go test -bench=. -cpu=1 -timeout=1h sort.go all_test.go
```

也可以参考 [912. 排序数组](https://leetcode.cn/problems/sort-an-array/)