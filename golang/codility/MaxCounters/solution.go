package MaxCounters

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Solution(N int, A []int) []int {
	maxN := 0
	base := 0
	answer := make([]int, N)
	for i := 0; i < len(A); i++ {
		idx := A[i]
		if idx == N+1 {
			base = maxN
		} else {
			if answer[idx-1] < base {
				answer[idx-1] = base
			}
			answer[idx-1]++
			maxN = max(maxN, answer[idx-1])
		}
	}
	for idx, v := range answer {
		if v < base {
			answer[idx] = base
		}
	}
	return answer
}
