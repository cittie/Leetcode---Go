package leetcode

// stack 确实是最佳解法……
func isValid(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}

	// 特别小的量级switch比map有优势
	getPair := func(left byte) byte {
		switch left {
		case '(':
			return ')'
		case '[':
			return ']'
		case '{':
			return '}'
		}
		return 0
	}

	stack := make([]byte, 0, n>>1)

	for i := 0; i < n; i++ {
		if c := getPair(s[i]); c != 0 {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != s[i] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

// 官方题解
func isValidOfficial(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []byte{}
	for i := 0; i < n; i++ {
		if pairs[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
