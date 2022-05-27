package mysort

import "math/bits"

func InsertionSort(arr []int) {
	l, r := 0, len(arr)
	for i := l + 1; i < r; i++ {
		for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
}

func QuickSort(arr []int) {
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
		QuickSort(arr[:p])
		QuickSort(arr[p+1:])
	}
}

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

func PDQSort(arr []int) {
	maxInsertion := 24
	if len(arr) <= maxInsertion {
		InsertionSort(arr)
		return
	}
	QuickSort(arr)
}
