package leetcode

import (
	"testing"
	"reflect"
	"sort"
)

func TestLetterCombinations(t *testing.T) {
	expected := []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}
	r := letterCombinations("23")

	sort.Strings(expected)
	sort.Strings(r)

	if !reflect.DeepEqual(expected, r) {
		t.Error("Expected: ", expected, "got: ", r)
	}
}
