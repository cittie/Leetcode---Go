package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombinationSum(t *testing.T) {
	tests := []struct {
		in     []int
		target int
		out    [][]int
	}{
		{[]int{2, 3, 6, 7}, 7, [][]int{{2, 2, 3}, {7}}},
		{[]int{2, 3, 5}, 8, [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}},
		{[]int{7, 3, 2}, 18, [][]int{
			{2, 2, 2, 2, 2, 2, 2, 2, 2}, {2, 2, 2, 2, 2, 2, 3, 3}, {2, 2, 2, 2, 3, 7}, {2, 2, 2, 3, 3, 3, 3}, {2, 2, 7, 7}, {2, 3, 3, 3, 7}, {3, 3, 3, 3, 3, 3}}},
	}

	for _, test := range tests {
		assert.Equal(t, test.out, combinationSum(test.in, test.target))
	}
}
