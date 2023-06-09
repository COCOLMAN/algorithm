package LongestPalindromicSubstring

import "fmt"

func check(s string) bool {
	fmt.Println(s)
	if len(s) == 1 {
		return true
	}
	if len(s) == 2 {
		return s[0] == s[1]
	}
	i := 0
	j := len(s) - 1

	if s[i] == s[j] {
		return check(s[1:len(s)-1]) || check(s[1:]) || check(s[:len(s)-1])
	}
	return false
}

func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}
	if check(s) {
		return s
	}
	return s
}
