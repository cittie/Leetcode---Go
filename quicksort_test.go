package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuicksort(t *testing.T) {
	tests := []struct {
		in  []int
		out []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{5, 1}, []int{1, 5}},
		{[]int{5, 5, 5}, []int{5, 5, 5}},
		{[]int{2, 1, 14, 25, 7, 2444}, []int{1, 2, 7, 14, 25, 2444}},
	}

	for _, test := range tests {
		qsort(test.in)
		assert.Equal(t, test.out, test.in)
	}
}
