package main

import "strings"

func lengthOfLastWordStrings(s string) int {
	trimmed := strings.TrimSpace(s)
	split := strings.Split(trimmed, " ")
	return len(split[len(split)-1])
}
