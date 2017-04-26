package _15__3Sum

import (
	"testing"
	"reflect"
	"sort"
)

func Test_threeSum(t *testing.T)  {
	nums := []int {-1, 0, 1, 2, -1, -4}
	expected := [][]int {
		[]int {-1, 0, 1},
		[]int {-1, -1, 2},
	}
	r := threeSum(nums)

	expected = twoDIntSliceSort(expected)
	r = twoDIntSliceSort(r)

	if !reflect.DeepEqual(r, expected) {
		t.Error("Expected: ", expected, "got: ", r,)
	}

}

func twoDIntSliceSort(ss [][]int) [][]int {

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