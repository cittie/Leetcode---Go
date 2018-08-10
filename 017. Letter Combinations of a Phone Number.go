package leetcode

/*
好像没什么可想的，空间换时间暴力遍历吧……
 */

var lettersMap = map[rune][]string {
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	var result []string
	for _, letter := range digits {
		result = genPermutation(result, lettersMap[letter])
	}
	return result
}

func genPermutation(a, b []string) []string {
	if len(a) == 0 {
		return b
	}

	if len(b) == 0 {
		return a
	}

	r := make([]string, 0, len(a)*len(b))
	for _, la := range a {
		for _, lb := range b {
			r = append(r, la+lb)
		}
	}

	return r
}
