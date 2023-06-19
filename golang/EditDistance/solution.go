package EditDistance

import "fmt"

var cache map[string]int

func min(candidate ...int) int {
	var m int
	m = candidate[0]
	for _, num := range candidate {
		if num < m {
			m = num
		}
	}
	return m
}

func getKey(k1, k2 string) string {
	return fmt.Sprintf("%s:%s", k1, k2)
}

func cal(word1, word2 string, op int) int {
	key := getKey(word1, word2)
	if val, exist := cache[key]; exist {
		return op + val
	}
	if len(word2) == 0 {
		return op + len(word1)
	}
	if len(word1) == 0 {
		return op + len(word2)
	}

	if word1[0] == word2[0] {
		val := min(
			cal(word1[1:], word2[1:], op),   // match
			cal(word1[1:], word2[1:], op+1), // replace
			cal(word1[1:], word2, op+1),     // delete
			cal(word1, word2[1:], op+1),     // insert
		)
		cache[key] = val - op
		return val
	}
	val := min(
		cal(word1[1:], word2[1:], op+1),
		cal(word1[1:], word2, op+1),
		cal(word1, word2[1:], op+1),
	)
	cache[key] = val - op
	return val
}

func minDistance(word1 string, word2 string) int {
	cache = map[string]int{}
	return cal(word1, word2, 0)
}
