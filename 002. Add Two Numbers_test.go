package leetcode

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	first := [][]int{
		[]int{2, 4, 3},
		[]int{9, 9, 9},
		[]int{3, 7, 9},
		[]int{1},
		[]int{1, 8},
	}
	second := [][]int{
		[]int{5, 6, 4},
		[]int{9, 9, 9},
		[]int{7, 2},
		[]int{9, 9, 9, 9},
		[]int{0},
	}
	expected := [][]int{
		[]int{7, 0, 8},
		[]int{8, 9, 9, 1},
		[]int{0, 0, 0, 1},
		[]int{0, 0, 0, 0, 1},
		[]int{1, 8},
	}

	for i := range first {
		p1, p2 := ConstructLinkedList(first[i]), ConstructLinkedList(second[i])
		p3 := addTwoNumbers(p1, p2)

		r := LinkedList2Slice(p3)

		if !reflect.DeepEqual(r, expected[i]) {
			t.Error("Expected: ", expected[i], "got: ", r)
		}
	}
}
