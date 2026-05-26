package main

var pair map[int]int = map[int]int{
	int(')'): int('('),
	int('}'): int('{'),
	int(']'): int('['),
}

func isValid(s string) bool {
	var parentheses []int = make([]int, 0, len(s))
	for _, char := range s {
		switch char {
		case '(', '{', '[':
			parentheses = append(parentheses, int(char))
		case ')', '}', ']':
			if len(parentheses) > 0 && parentheses[len(parentheses)-1] == pair[int(char)] {
				parentheses = parentheses[:len(parentheses)-1]
			} else {
				return false
			}
		default:
			return false
		}
	}
	return len(parentheses) == 0
}
