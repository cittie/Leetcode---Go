package leetcode

const (
	romanLen1 = 7
	romanLen2 = 6
)

var (
	nums1  = []int{1000, 500, 100, 50, 10, 5, 1}
	nums2  = []int{900, 400, 90, 40, 9, 4}
	rBytes = []byte{'M', 'D', 'C', 'L', 'X', 'V', 'I'}
	strs2  = []string{"CM", "CD", "XC", "XL", "IX", "IV"}

	m = map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	num int
)

func romanToInt(s string) int {
	var r int

	pre := 0
	for i := len(s) - 1; i >= 0; i-- {
		cur := m[s[i]]
		if cur >= pre {
			r += cur
		} else {
			r -= cur
		}
		pre = cur
	}

	return r
}

func romanToIntNew(s string) int {
	num = 0

	// slice长度较小时，遍历速度远远大于map
	// 所以选择空间换时间

loop:
	for len(s) > 0 {
		if len(s) >= 2 && (s[0] == 'C' || s[0] == 'X' || s[0] == 'I') {
			for i := 0; i < romanLen2; i++ {
				if s[:2] == strs2[i] {
					num += nums2[i]
					s = s[2:]
					continue loop
				}
			}
		}

		for i := 0; i < romanLen1; i++ {
			if s[0] == rBytes[i] {
				num += nums1[i]
				s = s[1:]
				break
			}
		}
	}

	return num
}
