package leetcode

import (
	"reflect"
	"testing"
)

func Test_partition(t *testing.T) {
	tests := []struct {
		in  []int
		x   int
		out []int
	}{
		{[]int{}, 15, []int{}},
		{[]int{1}, 2, []int{1}},
		{[]int{1, 2}, 0, []int{1, 2}},
		{[]int{1, 4, 3, 2, 5, 2}, 3, []int{1, 2, 2, 4, 3, 5}},
		{[]int{1, 1, 2, 2, 5, 6}, 2, []int{1, 1, 2, 2, 5, 6}},
	}

	for _, test := range tests {
		in := ConstructLinkedList(test.in)
		if r := partition(in, test.x); !reflect.DeepEqual(LinkedList2Slice(r), test.out) {
			t.Errorf("returned %v, want %v", LinkedList2Slice(r), test.out)
		}
	}
}
