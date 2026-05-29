package main

import "strings"

func wordPattern(pattern string, s string) bool {
	wordsInS := strings.Split(s, " ")
	metWords := make(map[string]bool, len(wordsInS))
	mapPattern := make(map[int32]string, len(wordsInS))
	if len(pattern) != len(wordsInS) {
		return false
	}
	for i, char := range pattern {
		word, ok := mapPattern[char]
		if ok {
			if word != wordsInS[i] {
				return false
			}
		} else {
			_, alreadyMet := metWords[wordsInS[i]]
			if alreadyMet {
				return false
			}
			mapPattern[char] = wordsInS[i]
			metWords[wordsInS[i]] = true
		}
	}
	return true
}
