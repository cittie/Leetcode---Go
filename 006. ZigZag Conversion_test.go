package leetcode

import "testing"

func Test_convert(t *testing.T) {
	tests := []struct {
		in  string
		n   int
		out string
	}{
		{"", 1, ""},
		{"abcd", 2, "acbd"},
		{"PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
		{"thisisatest", 4, "tahstiietss"},
	}

	for _, test := range tests {
		if r := convert(test.in, test.n); r != test.out {
			t.Errorf("returned %v, want %v", r, test.out)
		}
	}
}
