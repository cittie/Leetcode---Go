package leetcode

/*
	thought 2 costs more time when submit on leetcode,
	it will be better to use 2 for most of the time.
*/

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

func removeDuplicates2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	cur := 0

	for next := 1; next < len(nums); next++ {
		if nums[cur] != nums[next] {
			if cur+1 != next {
				nums[cur+1], nums[next] = nums[next], nums[cur+1]
			}
			cur++
		}
	}

	return cur + 1
}
