package leetcode

import (
	"strings"
)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	rowPos, step := 1, 1
	strSlice := make([]string, numRows)

	for _, c := range s {
		strSlice[rowPos-1] += string(c)
		// Reach top or bottom
		if rowPos%numRows == 0 || rowPos%numRows == 1 {
			step = -step
		}
		rowPos -= step
	}

	return strings.Join(strSlice, "")
}
