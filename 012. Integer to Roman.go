package leetcode

import "bytes"

const (
	romanLen = 13
)

var (
	nums = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	strs = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	buff bytes.Buffer
)

func intToRoman(num int) string {
	var roman string

	for i := 0; i < romanLen; i++ {
		for num >= nums[i] {
			roman += strs[i]
			num -= nums[i]
		}
	}

	return roman
}

func intToRomanNew(num int) string {
	buff.Reset()

	for i := 0; i < romanLen; i++ {
		for ; num >= nums[i]; num -= nums[i] {
			buff.WriteString(strs[i])
		}
	}

	return buff.String()
}
