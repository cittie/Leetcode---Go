package leetcode

import (
	"reflect"
	"testing"
)

func TestFourSum(t *testing.T) {
	nums := []int{-3, -2, -1, 0, 0, 1, 2, 3}
	target := 0
	expected := [][]int{
		{-3, -2, 2, 3},
		{-3, -1, 1, 3},
		{-3, 0, 0, 3},
		{-3, 0, 1, 2},
		{-2, -1, 0, 3},
		{-2, -1, 1, 2},
		{-2, 0, 0, 2},
		{-1, 0, 0, 1},
	}
	r := fourSum(nums, target)

	expected = TwoDIntSliceSort(expected)
	r = TwoDIntSliceSort(r)

	if !reflect.DeepEqual(r, expected) {
		t.Error("Expected: ", expected, "got: ", r)
	}
}

func Benchmark_fourSum(b *testing.B) {
	nums := []int{-3, -2, -1, 0, 0, 1, 2, 3}
	target := 0

	b.Run("old", func(b *testing.B) {
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			fourSum(nums, target)
		}
	})

	b.Run("official", func(b *testing.B) {
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			fourSumOfficial(nums, target)
		}
	})
}
