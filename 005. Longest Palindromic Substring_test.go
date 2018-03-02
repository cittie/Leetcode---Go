package leetcode

import "testing"

func Test_longestPalindrome(t *testing.T) {
	strings := []string{
		"babad",
		"bbbbb",
		"cbbd",
		"c",
		"bb",
		"aaa",
		"abcdefghijklmn",
		"abcda",
		"",
	}

	expected := []string{
		"bab",
		"bbbbb",
		"bb",
		"c",
		"bb",
		"aaa",
		"a",
		"a",
		"",
	}

	for i, s := range strings {
		if r := longestPalindrome(s); r != expected[i] {
			t.Error("Expected: ", expected[i], "got: ", r)
		}
	}
}
