package leetcode

import (
	"math"
)

func divide(dividend int, divisor int) int {
	isNegative := (dividend > 0) != (divisor > 0)

	dividendAbs, divisorAbs := dividend, divisor
	if dividend < 0 {
		dividendAbs = -dividend
	}
	if divisor < 0 {
		divisorAbs = -divisor
	}

	result := 0

	for divisorAbs <= dividendAbs {
		num, counts := divisorAbs, 1
		for num<<1 <= dividendAbs {
			num <<= 1
			counts <<= 1
		}
		dividendAbs -= num
		result += counts
	}

	if isNegative {
		result = -result
	}

	if result < math.MinInt32 || result > math.MaxInt32 {
		return math.MaxInt32
	}

	return result
}
