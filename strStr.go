package main

func strStr(haystack string, needle string) int {
	H := len(haystack)
	N := len(needle)
	for i := 0; i < H-N+1; i++ {
		if haystack[i:i+N] == needle {
			return i
		}
	}
	return -1
}
