package leetcode

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	reverted := 0
	for num := x; num > 0; num /= 10 {
		// reverted = reverted*10 + num%10		//
		reverted *= 10
		reverted += num % 10
	}

	return reverted == x
}
