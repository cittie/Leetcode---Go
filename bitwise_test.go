package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseInt(t *testing.T) {
	tests := []struct {
		in  int
		out int
	}{
		{0, 0},
		{1, -1},
		{427772, -427772},
		{-2234, 2234},
	}

	for _, test := range tests {
		assert.Equal(t, test.out, ReverseIntSign(test.in))
	}
}
