package leetcode

func maxArea(height []int) int {
	lower, area := 0, 0

	for left, right := 0, len(height)-1; left < right; {
		if lower = height[left]; height[right] < height[left] {
			lower = height[right]
		}

		if curArea := (right - left) * lower; curArea > area {
			area = curArea
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return area
}
