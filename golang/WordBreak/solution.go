package WordBreak

var cache map[string]bool

func search(maxLen int, wordMap *map[string]bool, s string) bool {
	if len(s) == 0 {
		return true
	}
	if val, exist := cache[s]; exist {
		return val
	}
	i := 0
	j := 1
	var result bool
	for j <= len(s) {
		target := s[i:j]
		if _, exist := (*wordMap)[target]; exist {
			if search(maxLen, wordMap, s[j:]) {
				result = true
				break
			}
		}

		if (j - i) >= maxLen {
			break
		}
		j++
	}
	cache[s] = result
	return result
}

func wordBreak(s string, wordDict []string) bool {
	cache = map[string]bool{}

	maxLen := 0

	wordMap := map[string]bool{}
	for _, word := range wordDict {
		if maxLen < len(word) {
			maxLen = len(word)
		}
		wordMap[word] = true
	}

	return search(maxLen, &wordMap, s)
}
