package main

import "fmt"

type Set = map[int]struct{}

func NewSet(slice []int) Set {
	set := make(Set)
	for _, v := range slice {
		set[v] = struct{}{}
	}
	return set
}

func main() {
	st := NewSet([]int{5, 4, 3, 2, 1})
	for i := range st {
		fmt.Println(i)
	}
}
