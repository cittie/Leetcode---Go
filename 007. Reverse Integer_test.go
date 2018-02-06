package leetcode

import (
	"testing"
)

func Test_reverse(t *testing.T) {
	nums := []int{2147483647, -2147483648, 2147483641, -2147483649, 2333, -982, 0}
	expectedNums := []int{0, 0, 1463847412, 0, 3332, -289, 0}

	for i, num := range nums {
		if r := reverse(num); r != expectedNums[i] {
			t.Error("Expected: ", expectedNums[i], "got: ", r)
		}
	}
}
