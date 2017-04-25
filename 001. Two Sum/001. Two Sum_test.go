package Leetcode___Go

import (
	"testing"
	"reflect"
	"sort"
)

func Test_TwoSum(t *testing.T)  {

	nums := []int{1, 2, 7, 11, 15}
	target := 9
	expected := []int{1, 2}

	r := twoSum(nums, target)

	sort.Ints(r)

	if !reflect.DeepEqual(r, expected) {
		t.Error("Expected: ", expected, "got: ", r,)
	}
}
