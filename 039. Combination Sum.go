package leetcode

import (
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	result := make([][]int, 0)
	initIntS := make([]int, 0)

	tryGetSum(target, 0, initIntS, &result, &candidates)

	return result
}

func tryGetSum(target, idx int, intS []int, results *[][]int, ca *[]int) {
	if len(*ca) == 0 {
		return
	}

	for i := idx; i < len(*ca); i++ {
		// make new slice instead using arg
		newIntS := make([]int, 0)
		newIntS = append(intS, (*ca)[i])

		if (*ca)[i] < target {
			if len(intS) != 0 && (*ca)[i] < intS[len(intS)-1] {
				continue
			}
			// same as above to avoid pointer value problem
			nextIntS := make([]int, len(newIntS))
			copy(nextIntS, newIntS)
			tryGetSum(target-(*ca)[i], i, nextIntS, results, ca)
		} else if (*ca)[i] == target {
			*results = append(*results, newIntS)
		} else {
			return
		}
	}
}
