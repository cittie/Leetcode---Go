package leetcode

func longestPalindrome(s string) string {
	n := len(s)
	result := ""

	check := func(left, right int) {
		for left >= 0 && right < n {
			if s[left] != s[right] {
				return
			}

			if len(s[left:right+1]) > len(result) {
				result = s[left : right+1]
			}

			left--
			right++
		}
	}

	for i := 0; i < len(s); i++ {
		check(i, i)
		check(i, i+1)
	}

	return result
}
