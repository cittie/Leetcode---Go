package leetcode

import (
	"math"
)

func myAtoi(s string) int {
	// cut space before
	idx := 0

	for i, ch := range s {
		if ch != rune(' ') {
			idx = i
			break
		}
	}

	if idx != 0 {
		s = s[idx:]
	}

	if s == "" {
		return 0
	}

	// copy Atoi from go src
	s0 := s
	if s[0] == '-' || s[0] == '+' {
		s = s[1:]
		if len(s) < 1 {
			return 0
		}
	}

	n := 0
	for _, ch := range []byte(s) {
		ch -= '0'
		if ch > 9 {
			break
		}
		n = n*10 + int(ch)
		if n < math.MinInt32 || n > math.MaxInt32 {
			break
		}
	}
	if s0[0] == '-' {
		n = -n
	}

	if n < math.MinInt32 {
		n = math.MinInt32
	} else if n > math.MaxInt32 {
		n = math.MaxInt32
	}

	return n
}
