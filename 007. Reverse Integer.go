package leetcode

func reverse(x int) int {
	var num, revNum int

	if num = x; x < 0 {
		num = ^x + 1
	}

	for num != 0 {
		revNum = 10*revNum + num%10
		num /= 10
		if revNum > (1<<31)-1 {
			return 0
		}
	}

	if x < 0 {
		revNum = ^revNum + 1
	}

	return revNum
}
