package leetcode

import (
	"testing"
)

func TestThreeSumClosest(t *testing.T) {
	tests := []struct {
		nums []int
		target int
		expeceted int
	}{
		{[]int{-1, 2, 1, -4}, 1, 2},
		{[]int{1,1,1,0}, -100, 2},
	}

	for _, test := range tests {
		if r := threeSumClosest(test.nums, test.target); r != test.expeceted {
			t.Error("Case:", test, " Expected:", test.expeceted, " got:", r)
		}
	}
}

func TestAbs(t *testing.T) {
	nums := []int64{-999, 0, 999}
	expected := []int64{999, 0, 999}

	for i, num := range nums {
		if abs(num) != expected[i] {
			t.Error("Expected: ", expected[i], "got: ", abs(num))
		}
	}
}