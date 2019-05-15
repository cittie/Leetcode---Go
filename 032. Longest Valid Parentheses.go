package leetcode

func longestValidParentheses(s string) int {
	result := 0
	newS := ")" + s
	dp := make([]int, len(newS))

	for i := 1; i < len(newS); i++ {
		if newS[i] == ')' {
			// matched
			if newS[i-1-dp[i-1]] == '(' {
				dp[i] = dp[i-1] + 2
			}

			// add previous longest
			dp[i] += dp[i-dp[i]]

			if dp[i] > result {
				result = dp[i]
			}
		}
	}

	return result
}
