package LongestPalindromicSubsequence

var cache map[string]int

func longestPalindromeSubseq(s string) int {
	if cache == nil {
		cache = map[string]int{}
	} else if val, exist := cache[s]; exist {
		return val
	}

	if len(s) < 2 {
		return len(s)
	}

	if s[0] == s[len(s)-1] {
		cache[s] = 2 + longestPalindromeSubseq(s[1:len(s)-1])
		return cache[s]
	}
	a := longestPalindromeSubseq(s[1:len(s)])
	b := longestPalindromeSubseq(s[0 : len(s)-1])
	if a > b {
		cache[s] = a
	} else {
		cache[s] = b
	}
	return cache[s]
}
