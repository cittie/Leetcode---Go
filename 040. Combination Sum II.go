package leetcode

import (
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	result := make([][]int, 0)
	initIntS := make([]int, 0)

	tryGetSum2(target, 0, initIntS, &result, &candidates)

	return result
}

func tryGetSum2(target, idx int, intS []int, results *[][]int, ca *[]int) {
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
			tryGetSum2(target-(*ca)[i], i+1, nextIntS, results, ca)
		} else if (*ca)[i] == target {
			// remove duplicates
			has := false
			for _, r := range *results {
				if equal(newIntS, r) {
					has = true
					break
				}
			}
			if !has {
				*results = append(*results, newIntS)
			}
		} else {
			return
		}
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
