package _03__Longest_Substring_Without_Repeating_Characters

func lengthOfLongestSubstring(s string) int {
	var last_positions = make(map[rune]int)
	var result, current int

	for i, c := range(s) {
		if last_length, ok := last_positions[c]; !ok || last_length < current {
			if i - current + 1 > result {
				result = i - current + 1
			}
		} else {
			current = last_positions[c] + 1
		}

		last_positions[c] = i
	}

	return  result
}