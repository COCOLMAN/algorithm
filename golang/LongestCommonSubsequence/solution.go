package LongestCommonSubsequence

import "fmt"

var cache map[string]int

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getKey(a, b string) string {
	return fmt.Sprintf("%s:%s", a, b)
}

func search(short, long string, size int) int {
	key := getKey(short, long)
	if val, exist := cache[key]; exist {
		return val + size
	}
	if len(short) == 0 || len(long) == 0 {
		return size
	}
	if short[0] == long[0] {
		a := search(short[1:], long[1:], size+1)
		cache[key] = a - size
		return cache[key] + size
	}

	a := max(
		search(short[1:], long, size),
		search(short, long[1:], size),
	)
	cache[key] = a - size
	return cache[key] + size
}

func longestCommonSubsequence(text1 string, text2 string) int {
	cache = map[string]int{}

	var long string = text1
	var short string = text2

	if len(text1) < len(text2) {
		long = text2
		short = text1
	}

	maxLen := 0
	for idx1, want := range short {
		for idx2, current := range long {
			if want == current {
				size := search(short[idx1:], long[idx2:], 0)
				if maxLen < size {
					maxLen = size
				}
			}
		}
	}
	/*
		oxcpqrsvwf
		shmtulqrypy
	*/
	return maxLen
}
