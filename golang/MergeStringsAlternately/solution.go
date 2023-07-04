package MergeStringsAlternately

import "strings"

func mergeAlternately(word1 string, word2 string) string {
	sb := strings.Builder{}
	maxLen := 0
	if len(word1) > len(word2) {
		maxLen = len(word1)
	} else {
		maxLen = len(word2)
	}
	for i := 0; i < maxLen; i++ {
		if i < len(word1) {
			sb.WriteByte(word1[i])
		}
		if i < len(word2) {
			sb.WriteByte(word2[i])
		}
	}
	return sb.String()
}
