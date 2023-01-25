package FrogRiverOne

func Solution(X int, A []int) int {
	fallHistory := map[int]int{}
	for at, position := range A {
		if _, ok := fallHistory[position]; !ok {
			fallHistory[position] = at
		}
	}

	x := 1
	maxAt := 0
	for x <= X {
		if at, ok := fallHistory[x]; !ok {
			return -1
		} else {
			if at > maxAt {
				maxAt = at
			}
		}
		x++
	}

	return maxAt
}
