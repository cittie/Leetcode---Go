package leetcode

import (
	"testing"
)

func Test_longestCommonPrefix(t *testing.T) {
	tests := []struct {
		in  []string
		out string
	}{
		{[]string{}, ""},
		{[]string{"", ""}, ""},
		{[]string{"one"}, "one"},
		{[]string{"one", "two"}, ""},
		{[]string{"one", "one"}, "one"},
		{[]string{"one", ""}, ""},
		{[]string{"one", "onForOne"}, "on"},
		{[]string{"one", "oneone", "oneoneone"}, "one"},
		{[]string{"oneoneone", "oneone", "one"}, "one"},
	}

	for _, test := range tests {
		if r := longestCommonPrefix(test.in); r != test.out {
			t.Errorf("%v returned %v, want %v", test.in, r, test.out)
		}
	}
}

func Benchmark_longestCommonPrefix(b *testing.B) {
	in := []string{"oneoneone", "oneone", "one"}

	b.Run("string", func(b *testing.B) {
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			longestCommonPrefix(in)
		}
	})

	b.Run("byte", func(b *testing.B) {
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			longestCommonPrefixNew(in)
		}
	})
}
