package MaxCounters

func getMaxNumber(A []int, m int) int {
	hash := map[int]int{}
	base := 0

	for i := len(A) - 1; i >= 0; i-- {
		n := A[i]
		if n == m {
			base += getMaxNumber(A[:i], m)
			break
		}
		if _, ok := hash[n]; !ok {
			hash[n] = 0
		}
		hash[n]++
	}
	maxN := 0
	for key := range hash {
		val := hash[key]
		if maxN < val {
			maxN = val
		}
	}
	return base + maxN
}

func Solution(N int, A []int) []int {
	var lastMaxCounterIdx int = 0
	var existMaxCounter = false

	for i := len(A) - 1; i >= 0; i-- {
		if A[i] == N+1 {
			lastMaxCounterIdx = i
			existMaxCounter = true
			break
		}
	}

	base := getMaxNumber(A[:lastMaxCounterIdx], N+1)
	answer := make([]int, N)
	for i := 0; i < N; i++ {
		answer[i] = base
	}

	if lastMaxCounterIdx == len(A)-1 {
		if existMaxCounter {
			return answer
		}
		lastMaxCounterIdx -= 1
	}
	for _, inc := range A[lastMaxCounterIdx+1:] {
		answer[inc-1]++
	}
	return answer
}
