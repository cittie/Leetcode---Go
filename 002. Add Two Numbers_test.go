package leetcode

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	first := [][]int{
		{2, 4, 3},
		{9, 9, 9},
		{3, 7, 9},
		{1},
		{1, 8},
	}
	second := [][]int{
		{5, 6, 4},
		{9, 9, 9},
		{7, 2},
		{9, 9, 9, 9},
		{0},
	}
	expected := [][]int{
		{7, 0, 8},
		{8, 9, 9, 1},
		{0, 0, 0, 1},
		{0, 0, 0, 0, 1},
		{1, 8},
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

func Benchmark_addTwoNumbers(b *testing.B) {
	p1, p2 := ConstructLinkedList([]int{9, 9, 9}), ConstructLinkedList([]int{9, 9, 9})

	for i := 0; i < b.N; i++ {
		addTwoNumbers(p1, p2)
	}
}
