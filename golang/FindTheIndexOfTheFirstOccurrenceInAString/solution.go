package FindTheIndexOfTheFirstOccurrenceInAString

func strStr(haystack string, needle string) int {
	i := 0
	candidate := []int{}
	for i < len(haystack)-len(needle)+1 {
		if haystack[i] == needle[0] {
			candidate = append(candidate, i)
		}
		i++
	}

	for _, c := range candidate {
		idx := c
		needleIdx := 0
		correct := true
		for idx < c+len(needle) {

			if haystack[idx] != needle[needleIdx] {
				correct = false
				break
			}
			idx++
			needleIdx++
		}
		if correct {
			return c
		}
	}
	return -1
}
