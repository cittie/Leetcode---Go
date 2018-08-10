package leetcode

import (
	"math"
	"sort"
)

/*
	思路：和015一样，排序后双指针遍历，记录当前的最接近值
	可以通过跳过相同的数字优化一点点，但是似乎不值得
 */

func threeSumClosest(nums []int, target int) int {
	result, minDistance := 0, math.MaxInt64
	sort.Ints(nums)
	n := len(nums)

	for cur:=0; cur < n-2; cur++ {
		left, right := cur+1, n-1

		/*
		// Skip duplicate current
		if cur > 0 && nums[cur] == nums[cur-1] {
			continue
		}
		*/

		for left < right {
			sum := nums[cur] + nums[left] + nums[right]
			curDis := int(abs(int64(sum - target)))

			// renew smallest
			if curDis < minDistance {
				minDistance = curDis
				result = sum
			}

			// move pointers
			if sum == target {
				break
			} else if sum > target {
				right--
			} else {
				left++
			}
		}
	}

	return result
}

func abs(n int64) int64 {
	y := n >> 63
	return (n ^ y) - y
}