package _03__Longest_Substring_Without_Repeating_Characters

import (
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T)  {
	strings := []string {
		"abcabcbb",
		"bbbbb",
		"pwwkew",
	}

	expected := []int {
		3,
		1,
		3,
	}

	for i, s := range(strings) {
		if r := lengthOfLongestSubstring(s); r != expected[i] {
			t.Error("Expected: ", expected[i], "got: ", r,)
		}

	}
}