package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	powerVal := 1
	for powerVal <= x {
		powerVal *= 10
	}
	topPower := powerVal / 10
	bottomPower := 1
	for topPower > bottomPower {
		leftDigit := (x / topPower) % 10
		rightDigit := (x / bottomPower) % 10
		if leftDigit != rightDigit {
			return false
		}
		topPower /= 10
		bottomPower *= 10
	}
	return true
}
