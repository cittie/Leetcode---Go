package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		in    []int
		out   int
		after []int
	}{
		{[]int{1, 1, 2}, 2, []int{1, 2}},
		{[]int{1}, 1, []int{1}},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5, []int{0, 1, 2, 3, 4}},
	}

	for _, test := range tests {
		assert.Equal(t, test.out, removeDuplicates(test.in))
		assert.Equal(t, test.after, test.in[:test.out])
	}
}

func BenchmarkRemoveDuplicates1(b *testing.B) {
	in := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	for i := 0; i < b.N; i++ {
		removeDuplicates(in)
	}
}

func BenchmarkRemoveDuplicates2(b *testing.B) {
	in := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	for i := 0; i < b.N; i++ {
		removeDuplicates2(in)
	}
}
