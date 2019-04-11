package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDivide(t *testing.T) {
	tests := []struct {
		dividend int
		divisor  int
		out      int
	}{
		{10, 3, 3},
		{7, -3, -2},
	}

	for _, test := range tests {
		assert.Equal(t, test.out, divide(test.dividend, test.divisor))
	}
}
