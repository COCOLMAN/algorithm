package PalindromePermutation

func canPermutePalindrome(s string) bool {
	counter := map[rune]int{}
	for _, c := range s {
		_, exist := counter[c]
		if !exist {
			counter[c] = 0
		}
		counter[c]++
	}

	oddCount := 0
	for _, count := range counter {
		if count%2 == 1 {
			oddCount++
			if oddCount > 1 {
				return false
			}
		}
	}
	return true
}
