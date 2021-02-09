package leetcode

import "bytes"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := ""

	for curPos := 0; curPos < len(strs[0]); curPos++ {
		curRune := strs[0][curPos]
		for pos := 1; pos < len(strs); pos++ {
			if len(strs[pos]) < curPos+1 || strs[pos][curPos] != curRune {
				return prefix
			}
		}

		prefix += string(curRune)
	}

	return prefix
}

func longestCommonPrefixNew(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	var buff bytes.Buffer

	for curPos := 0; curPos < len(strs[0]); curPos++ {
		curRune := strs[0][curPos]
		for pos := 1; pos < len(strs); pos++ {
			if len(strs[pos]) < curPos+1 || strs[pos][curPos] != curRune {
				return buff.String()
			}
		}

		buff.WriteByte(curRune)
	}

	return buff.String()
}
