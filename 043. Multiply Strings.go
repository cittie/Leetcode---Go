package leetcode

import (
	"strconv"
)

func multiply(num1 string, num2 string) string {
	n, m := len(num1), len(num2)

	// result max length is equal to num1 length plus num2 length
	result := make([]int, n+m)

	for i := n - 1; i >= 0; i-- {
		n1, _ := strconv.Atoi(string(num1[i]))
		for j := m - 1; j >= 0; j-- {
			n2, _ := strconv.Atoi(string(num2[j]))

			r := n1*n2 + result[i+j+1] // right side of current multiply result
			result[i+j+1] = r % 10
			result[i+j] += r / 10
		}
	}

	resultStr := ""
	for _, r := range result {
		if resultStr == "" && r == 0 {
			continue
		}
		resultStr += strconv.Itoa(r)
	}

	// oops miss this
	if resultStr == "" {
		resultStr = "0"
	}

	return resultStr
}
