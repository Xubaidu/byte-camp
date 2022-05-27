package main

import (
	"fmt"
	"hw9/mysort"
	"math/rand"
	"sort"
)

func Compare(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	var slice, stdSlice []int
	for i := 0; i < 10000; i++ {
		val := rand.Int()
		slice = append(slice, val)
		stdSlice = append(stdSlice, val)
	}
	sort.Ints(stdSlice)
	mysort.PDQSort(slice)
	result := Compare(slice, stdSlice)
	fmt.Println(result)
}
