package leetcode

import "testing"

func Test_findMedianSortedArrays(t *testing.T) {
	nums1 := [][]int{
		{1, 2},
		{1, 2},
	}
	nums2 := [][]int{
		{2},
		{3, 4},
	}

	expected := []float64{
		2.0,
		2.5,
	}

	for i := range nums1 {
		r := findMedianSortedArrays(nums1[i], nums2[i])
		if r != expected[i] {
			t.Error("Expected: ", expected[i], "got: ", r)
		}
	}
}

func Benchmark_findMedianSortedArrays(b *testing.B) {
	nums1, nums2 := []int{1, 2}, []int{3, 4}

	for i := 0; i < b.N; i++ {
		findMedianSortedArrays(nums1, nums2)
	}
}
