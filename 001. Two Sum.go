package leetcode

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	result := make([]int, 2)

	for i, num := range nums {
		_, ok := numMap[num]
		if ok {
			result[0], result[1] = i, numMap[num]
			break
		} else {
			numMap[target-num] = i
		}
	}

	return result
}

func twoSumOrg(nums []int, target int) []int {
	numMap := make(map[int]int)
	result := make([]int, 2)

	for i, num := range nums {
		_, ok := numMap[num]
		if ok {
			result[0], result[1] = i, numMap[num]
		} else {
			numMap[target-num] = i
		}
	}

	return result
}
