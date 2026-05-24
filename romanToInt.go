package main

func romanToInt(s string) int {
	result := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'M':
			result += 1000
		case 'D':
			result += 500
		case 'C':
			if i+1 < len(s) {
				switch s[i+1] {
				case 'D':
					result += 400
					i++
				case 'M':
					result += 900
					i++
				default:
					result += 100
				}
			} else {
				result += 100
			}
		case 'L':
			result += 50
		case 'X':
			if i+1 < len(s) {
				switch s[i+1] {
				case 'L':
					result += 40
					i++
				case 'C':
					result += 90
					i++
				default:
					result += 10
				}
			} else {
				result += 10
			}
		case 'V':
			result += 5
		case 'I':
			if i+1 < len(s) {
				switch s[i+1] {
				case 'V':
					result += 4
					i++
				case 'X':
					result += 9
					i++
				default:
					result += 1
				}
			} else {
				result += 1
			}
		}
	}
	return result
}
