package leetcode

import (
	"reflect"
	"testing"
)

func Test_threeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	expected := [][]int{
		[]int{-1, 0, 1},
		[]int{-1, -1, 2},
	}
	r := threeSum(nums)

	expected = TwoDIntSliceSort(expected)
	r = TwoDIntSliceSort(r)

	if !reflect.DeepEqual(r, expected) {
		t.Error("Expected: ", expected, "got: ", r)
	}

}
