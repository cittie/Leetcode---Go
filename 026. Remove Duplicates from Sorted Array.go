package leetcode

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	cur := 0

	for next := 1; next < len(nums); next++ {
		if nums[cur] != nums[next] {
			nums[cur+1], nums[next] = nums[next], nums[cur+1]
			cur++
		}
	}

	return cur + 1
}
