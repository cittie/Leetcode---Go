package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombinationSum2(t *testing.T) {
	tests := []struct {
		in     []int
		target int
		out    [][]int
	}{
		{[]int{10, 1, 2, 7, 6, 1, 5}, 8, [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}}},
		{[]int{2, 5, 2, 1, 2}, 5, [][]int{{1, 2, 2}, {5}}},
	}

	for _, test := range tests {
		assert.Equal(t, test.out, combinationSum2(test.in, test.target))
	}
}
