package PermCheck

func Solution(A []int) int {
	history := map[int]bool{}

	for _, a := range A {
		if _, exist := history[a]; exist {
			return 0
		}
		history[a] = true
	}

	N := len(A)
	n := 1

	for n <= N {
		_, exist := history[n]
		if !exist {
			return 0
		}
		n++
	}
	return 1
}
