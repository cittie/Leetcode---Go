package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_longestValidParentheses(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{"", 0},
		{"(()", 2},
		{")()())", 4},
		{")(()()))", 6},
	}

	for _, test := range tests {
		assert.Equal(t, test.out, longestValidParentheses(test.in))
	}
}

func Benchmark_longestValidParentheses(b *testing.B) {
	for i := 0; i < b.N; i++ {
		longestValidParentheses(")(()()))")
	}
}
