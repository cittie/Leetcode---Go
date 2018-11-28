package leetcode

import (
	"sort"
	"strconv"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)

	twoSums := make(map[int][]int)
	exists := make(map[string]struct{})
	results := make([][]int, 0)

	// get a map of sum of all 2 nums
	for i := 0; i < len(nums)-1; i++ {
		for j := i+1; j < len(nums); j++ {
			sum := nums[i] + nums[j]
			twoSums[sum] = append(twoSums[sum], i, j)
		}
	}

	// simply use "n1 n2 n3 n4 " as hash key
	nums2str := func(n []int) (numStr string) {
		for i := 0; i < len(n); i++ {
			numStr += strconv.Itoa(n[i]) + " "
		}
		return
	}

	// loop all 2sums
	for k, nums1 := range twoSums {
		if nums2, ok := twoSums[target - k]; ok {
			for i := 0; i < len(nums1); i+=2 {
				for j := 0; j < len(nums2); j+=2 {
					if nums1[i+1] < nums2[j] {
						curNum := []int{nums[nums1[i]], nums[nums1[i+1]], nums[nums2[j]], nums[nums2[j+1]]}
						// check duplicates
						if _, ok := exists[nums2str(curNum)]; !ok {
							results = append(results, curNum)
							exists[nums2str(curNum)] = struct{}{}
						}
					}
				}
			}
		} else {
			delete(twoSums, k)
		}
	}

	return results
}