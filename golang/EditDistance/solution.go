package EditDistance

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

func cal(word1, word2 string, op int) int {

	if len(word2) == 0 {
		return op + len(word1)
	}
	if len(word1) == 0 {
		return op + len(word2)
	}

	if word1[0] == word2[0] {
		return min(
			cal(word1[1:], word2[1:], op),   // match
			cal(word1[1:], word2[1:], op+1), // replace
			cal(word1[1:], word2, op+1),     // delete
			cal(word1, word2[1:], op+1),     // insert
		)
	}
	return min(
		cal(word1[1:], word2[1:], op+1),
		cal(word1[1:], word2, op+1),
		cal(word1, word2[1:], op+1),
	)
}

func minDistance(word1 string, word2 string) int {
	return cal(word1, word2, 0)
}
