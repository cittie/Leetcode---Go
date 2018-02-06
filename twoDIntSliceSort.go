package leetcode

import "sort"

// TwoDIntSliceSort returns a sorted 2D Int Slice
func TwoDIntSliceSort(ss [][]int) [][]int {

	// Sort each slice
	for _, s := range ss {
		sort.Sort(sort.IntSlice(s))
	}

	// Sort slices by each element within
	sort.Slice(ss, func(i, j int) bool {
		var n int

		// Set n to the length of the shorter slice
		if n = len(ss[i]); len(ss[j]) < len(ss[i]) {
			n = len(ss[j])
		}

		// Compare the first nth elements in slices
		for k := 0; k < n; k++ {
			if ss[i][k] != ss[j][k] {
				return ss[i][k] < ss[j][k]
			}
		}

		// Return whatever if the 2 slices are the same.
		return true
	})

	return ss
}
