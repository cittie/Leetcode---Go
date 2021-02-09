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
		for j := i + 1; j < len(nums); j++ {
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
		if nums2, ok := twoSums[target-k]; ok {
			for i := 0; i < len(nums1); i += 2 {
				for j := 0; j < len(nums2); j += 2 {
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

// 官方解法
// 果然快好多……
// 剪枝NB!
func fourSumOfficial(nums []int, target int) (quadruplets [][]int) {
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n-3 && nums[i]+nums[i+1]+nums[i+2]+nums[i+3] <= target; i++ {
		// 剪枝
		if i > 0 && nums[i] == nums[i-1] || nums[i]+nums[n-3]+nums[n-2]+nums[n-1] < target {
			continue
		}
		for j := i + 1; j < n-2 && nums[i]+nums[j]+nums[j+1]+nums[j+2] <= target; j++ {
			// 剪枝
			if j > i+1 && nums[j] == nums[j-1] || nums[i]+nums[j]+nums[n-2]+nums[n-1] < target {
				continue
			}
			// 双指针
			for left, right := j+1, n-1; left < right; {
				if sum := nums[i] + nums[j] + nums[left] + nums[right]; sum == target {
					quadruplets = append(quadruplets, []int{nums[i], nums[j], nums[left], nums[right]})
					for left++; left < right && nums[left] == nums[left-1]; left++ {
					}
					for right--; left < right && nums[right] == nums[right+1]; right-- {
					}
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}
	return
}
