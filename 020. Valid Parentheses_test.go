package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isValid(t *testing.T) {
	tests := []struct {
		s string
		r bool
	}{
		{"", true},
		{"(", false},
		{"()", true},
	}

	for i, test := range tests {
		assert.Equal(t, test.r, isValid(test.s), i)
	}
}

func Benchmark_isValid(b *testing.B) {
	b.Run("old", func(b *testing.B) {
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			isValid("([)([)([)]]([)](([)]([)][([)]([)])]]")
		}
	})

	b.Run("official", func(b *testing.B) {
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			isValidOfficial("([)([)([)]]([)](([)]([)][([)]([)])]]")
		}
	})
}
