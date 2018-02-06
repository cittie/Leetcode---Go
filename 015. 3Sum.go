package leetcode

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	results := [][]int{}
	sort.Ints(nums)
	n := len(nums)

	for cur := 0; cur < n-2; cur++ {
		left, right := cur+1, n-1

		// Skip duplicate current
		if cur > 0 && nums[cur] == nums[cur-1] {
			continue
		}

		for left < right {
			total := nums[cur] + nums[left] + nums[right]

			if total == 0 {
				results = append(results, []int{nums[cur], nums[left], nums[right]})

				// Skip duplicate left
				for right > left && nums[left] == nums[left+1] {
					left++
				}

				// Skip duplicate right
				for right > left && nums[right] == nums[right-1] {
					right--
				}

				left++
				right--
			} else if total < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return results
}
