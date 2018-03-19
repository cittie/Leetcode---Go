package leetcode

import (
	"math"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		in  int
		out bool
	}{
		{0, true},
		{11, true},
		{515, true},
		{1, true},
		{101, true},
		{3333, true},
		{32, false},
		{123, false},
		{123322221, false},
		{-222, false},
		{math.MaxInt64, false},
	}

	for _, test := range tests {
		if r := isPalindrome(test.in); r != test.out {
			t.Errorf("%v returned %v, want %v", test.in, r, test.out)
		}
	}
}
