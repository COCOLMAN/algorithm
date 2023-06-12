package LongestPalindromicSubstring

func searchPalindromeFromIdx(s string, left, right int) string {
	if len(s) == 0 {
		return ""
	}
	diff := 0
	maxPalindrome := s[0:1]
	for left-diff >= 0 && right+diff <= len(s)-1 {
		if s[left-diff] != s[right+diff] {
			break
		}
		maxPalindrome = s[left-diff : right+diff+1]
		diff++
	}
	return maxPalindrome
}

func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}
	if len(s) == 2 {
		if s[0] == s[1] {
			return s
		} else {
			return s[0:1]
		}
	}
	maxPalindrome := ""
	for i := 0; i <= len(s)-2; i++ {
		currentPalindrome1 := searchPalindromeFromIdx(s, i, i+1)
		currentPalindrome2 := searchPalindromeFromIdx(s, i-1, i+1)
		currentPalindrome := ""
		if len(currentPalindrome1) > len(currentPalindrome2) {
			currentPalindrome = currentPalindrome1
		} else {
			currentPalindrome = currentPalindrome2
		}
		if len(currentPalindrome) > len(maxPalindrome) {
			maxPalindrome = currentPalindrome
		}
	}

	return maxPalindrome
}
