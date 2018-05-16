package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxArea(t *testing.T) {
	tests := []struct {
		in  []int
		out int
	}{
		{[]int{2, 2}, 2},
		{[]int{2, 3, 4}, 4},
		{[]int{1, 2, 3, 4, 3, 4, 2, 1}, 10},
		{[]int{5, 4, 3, 2, 1}, 6},
	}

	for _, test := range tests {
		assert.Equal(t, test.out, maxArea(test.in))
	}
}
