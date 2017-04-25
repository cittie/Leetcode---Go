package Leetcode___Go

func twoSum(nums []int, target int) []int {
	numMap := map[int]int{}
	result := make([]int, 2)

	for i := range nums {
		_, ok := numMap[nums[i]]
		if ok {
			result[0] = i
			result[1] = numMap[nums[i]]
		} else {
			numMap[target-nums[i]] = i
		}
	}

	return result
}
