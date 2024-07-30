package LongestSubstringWithAtMostTwoDistinctCharacters

import "fmt"

func lengthOfLongestSubstringTwoDistinct(s string) int {
	i := 0
	ch := s[i]
	counts := []int{1}
	for i < len(s)-1 {
		i++

		if s[i] == ch {
			counts[len(counts)-1]++
		} else {
			counts = append(counts, 1)
			ch = s[i]
		}
	}
	fmt.Println(counts)
	return 0
}
