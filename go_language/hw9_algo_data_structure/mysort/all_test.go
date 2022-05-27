package mysort

import (
	"math/rand"
	"sort"
	"testing"
)

func TestAll(t *testing.T) {
	fuzzTestSort(t, func(data []int) {
		InsertionSort(data)
	})
	fuzzTestSort(t, func(data []int) {
		HeapSort(data)
	})
	fuzzTestSort(t, func(data []int) {
		QuickSort(data)
	})
	fuzzTestSort(t, func(data []int) {
		PDQSort(data)
	})
}

func fuzzTestSort(t *testing.T, f func(data []int)) {
	const times = 2048
	randomTestTimes := rand.Intn(times)
	for i := 0; i < randomTestTimes; i++ {
		randomLenth := rand.Intn(times)
		v1 := make([]int, randomLenth)
		v2 := make([]int, randomLenth)
		for j := 0; j < randomLenth; j++ {
			randomValue := rand.Intn(randomLenth)
			v1[j] = randomValue
			v2[j] = randomValue
		}
		sort.Ints(v1)
		f(v2)
		for idx := range v1 {
			if v1[idx] != v2[idx] {
				t.Fatal("invalid sort:", idx, v1[idx], v2[idx])
			}
		}
	}
}
