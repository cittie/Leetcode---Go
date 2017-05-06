package _04__Median_of_Two_Sorted_Array

import "testing"

func Test_findMedianSortedArrays(t *testing.T)  {
	nums1 :=[][]int{
		[]int {1, 2},
		[]int {1, 2},
	}
	nums2 := [][]int {
		[]int {2},
		[]int {3, 4},
	}

	expected := []float64 {
		2.0,
		2.5,
	}

	for i := range(nums1) {
		r := findMedianSortedArrays(nums1[i], nums2[i])
		if r != expected[i] {
			t.Error("Expected: ", expected[i], "got: ", r,)
		}
	}

}