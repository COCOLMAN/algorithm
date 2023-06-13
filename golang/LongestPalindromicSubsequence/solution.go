package LongestPalindromicSubsequence

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func search(s string, left, right, size int) int {

	if left < 0 || right >= len(s) {
		return size
	}
	if s[left] == s[right] {
		size++
	}

	result := max(
		search(s, left-1, right, size),
		search(s, left, right+1, size),
	)
	return result
}

func longestPalindromeSubseq(s string) int {
	maxVal := 0
	fmt.Println(s, s[:len(s)-1])
	for idx, _ := range s[:len(s)-1] {

		var a, b int
		a = search(s, idx, idx, 0)
		fmt.Println(idx, idx, ":", a)
		if s[idx] == s[idx+1] {
			b = search(s, idx, idx+1, 1)
			fmt.Println(idx, idx+1, b)
		}

		val := max(a, b)
		if val > maxVal {
			maxVal = val
		}
		fmt.Println(s, idx, val)
	}
	return maxVal
}
