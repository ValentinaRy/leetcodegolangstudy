package main

func longestCommonPrefix(strs []string) string {
	minLen := len(strs[0])
	for _, str := range strs {
		if len(str) < minLen {
			minLen = len(str)
		}
	}
	for i := 0; i < minLen; i++ {
		curLetter := strs[0][i]
		for _, str := range strs {
			if str[i] != curLetter {
				return strs[0][:i]
			}
		}
	}
	return strs[0][:minLen]
}
