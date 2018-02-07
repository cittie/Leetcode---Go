package leetcode

func integerBreak(n int) int {
	switch n {
	case 2:
		return 1
	case 3:
		return 2
	}

	total := 0

	intPow := func(x, y int) int {
		result := 1

		for y > 0 {
			if y&1 != 0 {
				result *= x
			}
			y >>= 1
			x *= x
		}

		return result
	}

	switch n % 3 {
	case 0:
		total = intPow(3, (n / 3))
	case 1:
		total = intPow(3, (n-4)/3) * 4
	case 2:
		total = intPow(3, (n-2)/3) * 2
	}

	return total
}

func integerBreak2(n int) int {
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}

	total := 1

	for n > 4 {
		total *= 3
		n -= 3
	}

	total *= n

	return total
}
