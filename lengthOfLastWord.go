package main

func lengthOfLastWord(s string) int {
	var i int = len(s) - 1
	for i >= 0 {
		if s[i] != ' ' {
			break
		}
		i--
	}
	wordLength := 0
	for i >= 0 {
		if s[i] == ' ' {
			break
		}
		i--
		wordLength += 1
	}
	return wordLength
}
