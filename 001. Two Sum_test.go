package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

func Test_twoSum(t *testing.T) {

	nums := []int{1, 2, 7, 11, 15}
	target := 9
	expected := []int{1, 2}

	r := twoSum(nums, target)

	sort.Ints(r)

	if !reflect.DeepEqual(r, expected) {
		t.Error("Expected: ", expected, "got: ", r)
	}
}
