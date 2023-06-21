package MinimumASCIIDeleteSumForTwoStrings

import "fmt"

var cache map[string]int

func ascii(a uint8) int {
	return int(a)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func getASCIISum(s string) int {
	total := 0
	for _, c := range s {
		total += ascii(uint8(c))
	}
	return total
}

func getKey(s1, s2 string) string {
	return fmt.Sprintf("%s:%s", s1, s2)
}

func cal(s1, s2 string, total int) int {
	key := getKey(s1, s2)
	if val, exist := cache[key]; exist {
		return val + total
	}
	if len(s1) == 0 {
		return total + getASCIISum(s2)
	}
	if len(s2) == 0 {
		return total + getASCIISum(s1)
	}
	if s1[0] == s2[0] {
		result := cal(s1[1:], s2[1:], total)
		cache[key] = result - total
		return cache[key] + total
	}
	result := min(
		cal(s1[1:], s2, total+ascii(s1[0])),
		cal(s1, s2[1:], total+ascii(s2[0])),
	)
	cache[key] = result - total
	return cache[key] + total
}

func minimumDeleteSum(s1 string, s2 string) int {
	cache = map[string]int{}
	return cal(s1, s2, 0)
}
