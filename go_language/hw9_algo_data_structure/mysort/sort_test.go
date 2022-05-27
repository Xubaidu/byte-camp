package mysort

import (
	"testing"
)

type args struct {
	arr []int
}

type testCase struct {
	name string
	args args
	want []int
}

var testCases = []testCase{
	{
		name: "test 1",
		args: args{
			arr: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		},
		want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	},
	{
		name: "test 2",
		args: args{
			arr: []int{6, 6, 6, 6, 6, 5, 5, 5, 5, 4, 4, 4},
		},
		want: []int{4, 4, 4, 5, 5, 5, 5, 6, 6, 6, 6, 6},
	},
	{
		name: "test 3",
		args: args{
			arr: []int{3, 5, 6, 2, 1, 7},
		},
		want: []int{1, 2, 3, 5, 6, 7},
	},
}

func Check(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

func TestInsertionSort(t *testing.T) {
	tests := testCases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if InsertionSort(tt.args.arr); !Check(tt.args.arr, tt.want) {
				t.Errorf("solve() = %v, want %v", tt.args.arr, tt.want)
			}
		})
	}
}

func TestQuickSort(t *testing.T) {
	tests := testCases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if QuickSort(tt.args.arr); !Check(tt.args.arr, tt.want) {
				t.Errorf("solve() = %v, want %v", tt.args.arr, tt.want)
			}
		})
	}
}

func TestHeapSort(t *testing.T) {
	tests := testCases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if HeapSort(tt.args.arr); !Check(tt.args.arr, tt.want) {
				t.Errorf("solve() = %v, want %v", tt.args.arr, tt.want)
			}
		})
	}
}

func TestPDQSort(t *testing.T) {
	tests := testCases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if PDQSort(tt.args.arr); !Check(tt.args.arr, tt.want) {
				t.Errorf("solve() = %v, want %v", tt.args.arr, tt.want)
			}
		})
	}
}
