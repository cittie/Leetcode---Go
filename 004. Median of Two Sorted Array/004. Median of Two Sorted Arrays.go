package _04__Median_of_Two_Sorted_Array

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n := (len(nums1) + len(nums2))		// Used to check odd or even
	median := (n + 1) >> 1
	var r float64

	if n & 1 != 0 {
		r = float64(bisearch(nums1, nums2, median))
	} else {
		r = float64((bisearch(nums1, nums2, median) +
			bisearch(nums1, nums2, median + 1))) / float64(2)
	}

	return r
}

func bisearch(nums1 []int, nums2 []int, median int) int {
	n1, n2 := len(nums1), len(nums2)

	if n1 < n2 {
		return bisearch(nums2, nums1, median)		// Reverse
	} else if n2 == 0 {
		return nums1[median - 1]
	} else if median == 1 {
		val := nums1[0]
		if nums1[0] > nums2[0] {
			val = nums2[0]
		}
		return val
	}

	median2 := n2
	if (median >> 1) < n2 {
		median2 = (median >> 1 )
	}
	median1 := median - median2

	if nums1[median1 - 1] < nums2[median2 - 1] {
		return bisearch(nums1[median1:], nums2, median - median1)
	} else if nums1[median1 - 1] > nums2[median2 - 1] {
		return bisearch(nums1, nums2[median2:], median - median2)
	} else {
		return nums1[median1 - 1]
	}
}