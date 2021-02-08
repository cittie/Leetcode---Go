package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_myAtoi(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{"", 0},
		{"42", 42},
		{"   -42", -42},
		{"4193 with words", 4193},
		{"words and 987", 0},
		{"-91283472332", -2147483648},
		{"-912834723322323222356787978780789676575685685686756756756756", -2147483648},
	}

	for _, test := range tests {
		assert.Equal(t, test.out, myAtoi(test.in))
	}
}

func BenchmarkMyAtoi(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		myAtoi("-912834723322323222356787978780789676575685685686756756756756")
	}
}
