package LengthOfLastWord

func lengthOfLastWord(s string) int {
	words := []string{}
	word := ""
	for _, c := range s {
		if string(c) == " " {
			if len(word) != 0 {
				words = append(words, word)
			}
			word = ""
		} else {
			word = word + string(c)
		}
	}
	if len(word) != 0 {
		words = append(words, word)
	}
	return len(words[len(words)-1])
}
