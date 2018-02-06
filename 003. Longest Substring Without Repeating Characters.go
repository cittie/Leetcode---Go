package leetcode

func lengthOfLongestSubstring(s string) int {
	lastPositions := make(map[rune]int)
	var result, current int

	for i, c := range s {
		if lastLength, ok := lastPositions[c]; !ok || lastLength < current {
			if curLength := i - current + 1; curLength > result {
				result = curLength
			}
		} else {
			current = lastPositions[c] + 1
		}

		lastPositions[c] = i
	}

	return result
}
