package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeNthFromEnd(t *testing.T) {
	tests := []struct {
		in    []int
		n     int
		after []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{1, 2, 3, 5}},
		{[]int{1}, 1, []int{}},
		{[]int{1, 2}, 1, []int{1}},
		{[]int{1, 2, 3, 4, 5}, 5, []int{2, 3, 4, 5}},
	}

	for _, test := range tests {
		n := removeNthFromEnd(ConstructLinkedList(test.in), test.n)
		assert.Equal(t, test.after, LinkedList2Slice(n))
	}
}
