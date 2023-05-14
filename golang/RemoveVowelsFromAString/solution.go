package RemoveVowelsFromAString

import "strings"

func removeVowels(s string) string {
	var newS strings.Builder
	vowels := map[rune]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}

	for _, c := range s {
		if _, exist := vowels[c]; !exist {
			newS.WriteRune(c)
		}
	}

	return newS.String()
}
