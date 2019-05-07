package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_multiply(t *testing.T) {
	tests := []struct {
		num1     string
		num2     string
		expected string
	}{
		{"2", "3", "6"},
		{"123", "456", "56088"},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, multiply(test.num1, test.num2))
	}
}
