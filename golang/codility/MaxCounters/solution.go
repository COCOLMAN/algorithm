package MaxCounters

func max(A []int) int {
	maxN := A[0]
	for _, n := range A {
		if n > maxN {
			maxN = n
		}
	}
	return maxN
}

func solution(N int, A []int, base int) []int {
	list := make([]int, N)
	for i := 0; i < N; i++ {
		list[i] = base
	}
	for idx, a := range A {
		if a == N+1 {
			return solution(N, A[idx+1:], max(list))
		}
		list[a-1]++
	}
	return list
}

func Solution(N int, A []int) []int {
	return solution(N, A, 0)
}
