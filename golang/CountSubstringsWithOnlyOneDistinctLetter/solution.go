package CountSubstringsWithOnlyOneDistinctLetter

func getSum(n int) int {
	return n * (n + 1) / 2
}

func countLetters(s string) int {
	total := 0
	distinctLength := 1
	currentLetter := rune(s[0])
	for _, char := range s[1:] {
		if char == currentLetter {
			distinctLength++
		} else {
			total += getSum(distinctLength)
			distinctLength = 1
			currentLetter = char
		}
	}
	total += getSum(distinctLength)
	return total
}
