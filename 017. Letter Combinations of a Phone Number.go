package leetcode

/*
  好像没什么可想的，空间换时间暴力遍历吧……
*/

var lettersMap = map[rune][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"p", "q", "r", "s"},
	'9': {"w", "x", "y", "z"},
}

var letters = [][]string{
	{},
	{},
	{"a", "b", "c"},      // 2
	{"d", "e", "f"},      // 3
	{"g", "h", "i"},      // 4
	{"j", "k", "l"},      // 5
	{"m", "n", "o"},      // 6
	{"p", "q", "r", "s"}, // 7
	{"p", "q", "r", "s"}, // 8
	{"w", "x", "y", "z"}, // 9
}

var result []string

func letterCombinations(digits string) []string {
	result = result[:0]
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

// dfs
func letterCombinationsNew(digits string) []string {
	result = result[:0]
	var keys []string = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	if digits == "" {
		return result
	}
	dfs(digits, keys, 0, "", &result)
	return result
}

func dfs(graph string, colors []string, index int, path string, res *[]string) {
	if index == len(graph) {
		*res = append(*res, path)
		return
	}
	actions := int(graph[index] - '0')
	for _, v := range colors[actions] {
		dfs(graph, colors, index+1, path+string(v), res)
	}
}
