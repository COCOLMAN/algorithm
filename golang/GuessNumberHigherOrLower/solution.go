package GuessNumberHigherOrLower

func generateGuess(pick int) func(int) int {
	guess := func(n int) int {
		if pick == n {
			return 0
		} else if pick < n {
			return -1
		} else {
			return 1
		}
	}
	return func(n int) int {
		start, end := 1, n
		for start <= end {
			mid := (start + end) / 2
			answer := guess(mid)
			if answer == 0 {
				return mid
			} else if answer == -1 {
				end = mid - 1
			} else {
				start = mid + 1
			}
		}
		return -1
	}
}
