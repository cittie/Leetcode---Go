package leetcode

// ReverseIntSign 将a转换为-a
func ReverseIntSign(a int) int {
	return ^a + 1
}
