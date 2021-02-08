package leetcode

func isMatch(s string, p string) bool {
	lenS, lenP := len(s), len(p)

	// p为空
	if lenP == 0 {
		return lenS == 0
	}

	if lenP > 1 && p[1] == '*' {
		return (lenS > 0 && (s[0] == p[0] || p[0] == '.')) &&
			isMatch(s[1:], p) || isMatch(s, p[2:])
	}

	return lenS > 0 &&
		(s[0] == p[0] || p[0] == '.') &&
		isMatch(s[1:], p[1:])
}

func isMatchDP(s string, p string) bool {
	lenS, lenP := len(s), len(p)

	// dp[i][j] = s[:i] == p[:j]
	dp := make([][]bool, lenS+1)
	for i := range dp {
		dp[i] = make([]bool, lenP+1)
	}

	dp[0][0] = true

	/*
		情况总结：
		1. dp[i][j] = dp[i-1][j-1]
			- p[j-1] != '*'
				- p[j-1] == s[i-1]
				- p[j-1] == '.'
		2. dp[i][j] = dp[i][j-2]
			- p[j-1] == '*'
			- s对应字符出现0次
		3. dp[i][j] = dp[i-1][j]
			- p[j-1] == '*'
				- p[j-2] == s[i-1]
				- p[j-2] == '.'
	*/

	for i := 0; i <= lenS; i++ {
		for j := 1; j <= lenP; j++ {
			if j > 1 && p[j-1] == '*' {
				dp[i][j] = dp[i][j-2] ||
					(i > 0 && (p[j-2] == s[i-1] || p[j-2] == '.') && dp[i-1][j])
			} else {
				dp[i][j] = i > 0 &&
					dp[i-1][j-1] &&
					(p[j-1] == s[i-1] || p[j-1] == '.')
			}
		}
	}

	return dp[lenS][lenP]
}
