# 排序算法

## 稳定排序的概念

对于值相同的元素，如果排序前后不改变它们的相对位置，那就是稳定排序，否则是不稳定排序。

稳定排序的应用场景主要是多关键词排序，可以减少排序的比较次数。

考虑如下两列数据，使用非稳定排序可能会打乱 `2(1), 2(2)` 的位置，从而还需要对第二列排序，所以比较次数更多。

```cpp
3       1
2(1)    4
2(2)    5
6       7
```

## 插入排序

最差时间复杂度 $\mathcal{O}(n^2)$，发生在完全逆序的情况

最有时间复杂度 $\mathcal{O}(n)$，发生在顺序的情况

插入排序是稳定排序算法，缓存友好（顺序取数）

```go
// go
func InsertionSort(arr []int) {
    l, r := 0, len(arr)
    for i := l + 1; i < r; i++ {
        for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
            arr[j], arr[j-1] = arr[j-1], arr[j]
        }
    }
}
```

## 堆排序

建堆时间复杂度为 $\mathcal{O}(n)$，具体来讲，设 $H=\log_2n$
$$
\sum\limits_{i = 0}^{H-1}(H-i-1)\cdot 2^{i} = \mathcal{O}(n)
$$

堆排序复杂度为 $\mathcal{O}(n\log n)$

堆排序缓存不友好（跳跃着取数）

堆排序不稳定（考虑第一个数字和第二个数字构成了 top2，第一次执行完毕后第一个数字会被交换到数组尾部）

```go
// go
func heapify(arr []int, i int) {
    left := 2*i + 1
    right := 2*i + 2
    largest := i
    if left < len(arr) && arr[left] > arr[largest] {
        largest = left
    }
    if right < len(arr) && arr[right] > arr[largest] {
        largest = right
    }
    if largest != i {
        arr[i], arr[largest] = arr[largest], arr[i]
        heapify(arr, largest)
    }
}

func HeapSort(arr []int) {
    BuildMaxHeap := func(arr []int) {
        for i := (len(arr) - 1) / 2; i >= 0; i-- {
            heapify(arr, i)
        }
    }
    BuildMaxHeap(arr)
    for i := len(arr) - 1; i > 0; i-- {
        arr[0], arr[i] = arr[i], arr[0]
        heapify(arr[:i], 0)
    }
}
```

## 快速排序

快速排序期望时间复杂度为 $\mathcal{O}(n\log n)$，出现在每次都取中位数的情况

最坏为 $\mathcal{O}(n^2)$，出现在每次都分割成 `1 : n-1` 的情况

快速排序缓存相对友好，是不稳定排序

```go
// go
func QuickSortForPDQ(arr []int) {
    partition := func(arr []int) int {
        pivot := arr[0]
        i, j := 0, len(arr)-1
        for i < j {
            for i < j && arr[j] >= pivot {
                j--
            }
            arr[i] = arr[j]
            for i < j && arr[i] <= pivot {
                i++
            }
            arr[j] = arr[i]
        }
        arr[i] = pivot
        return i
    }
    if len(arr) > 1 {
        p := partition(arr)
        leftDist, rightDist := p+1, len(arr)-p
        limit := bits.Len(uint(len(arr)))
        if p <= leftDist || p <= rightDist {
            limit--
        }
        if limit == 0 {
            HeapSort(arr)
            return
        }
        QuickSort(arr[:p])
        QuickSort(arr[p+1:])
    }
}
```

## PDQsort

`pattern defeating quicksort`，结合了上述三种排序，互相补充，保障了最坏复杂度为 $\mathcal{O}(n
\log n)$

### 由实验获得的一些启发

- 所有短序列和有序情况下，`InsertionSort` 效果最好
- 大部分情况下，`QuickSort` 效果最好，或许是因为缓存友好的原因
- `HeapSort` 用来兜底，保障最差复杂度为 $\mathcal{O}(n\log n)$

### PDQSort 细节

- 数组长度不超过 `24` 时，使用 `InsertionSort`（泛型情况下，这个值最优秀）
- 设定 `limit = bits.Len(length)`，如果 `pivot` 距离两端点太近，`i.e.` 不超过 `length/8` 时，`limit` 减一，当 `limit` 为 `0` 时，使用 `HeapSort` 保障复杂度
- 一般情况使用 `QuickSort`

### 优化手段

- 升级 pivot 选择策略（近似中位数）
- 采样
  - 如果可能逆序，则翻转数组，应对逆序情况
  - 如果可能顺序，直接插排