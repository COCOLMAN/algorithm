package LongestPalindromicSubstring

func check(s string, cache *map[string]bool) bool {
	if v, exist := (*cache)[s]; exist {
		return v
	}
	if len(s) == 1 {
		return true
	}
	if len(s) <= 3 {
		return s[0] == s[len(s)-1]
	}
	i := 0
	j := len(s) - 1

	for i < j {
		if s[i] != s[j] {
			(*cache)[s] = false
			return false
		}
		i++
		j--
	}
	(*cache)[s] = true
	return true
}

func ggg(s string, cache *map[string]bool, cache2 *map[string]string) string {
	if v, exist := (*cache2)[s]; exist {
		return v
	}
	i := 0
	j := len(s)

	if check(s[i:j], cache) {
		r := s[i:j]
		(*cache2)[s] = r
		return r
	}
	a := ggg(s[i+1:j], cache, cache2)
	b := ggg(s[i:j-1], cache, cache2)
	if len(a) > len(b) {
		(*cache2)[s] = a
		return a
	}
	(*cache2)[s] = b
	return b
}

func longestPalindrome(s string) string {
	cache := map[string]bool{}
	cache2 := map[string]string{}
	return ggg(s, &cache, &cache2)
}
