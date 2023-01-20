package MaxCounters

func max(d map[int]int) int {
	maxN := 0
	for key := range d {
		if d[key] > maxN {
			maxN = d[key]
		}
	}
	return maxN
}

func get(N int, A []int) int {
	hash := map[int]int{}
	for idx, a := range A {
		if a == N+1 {
			return max(hash) + get(N, A[idx+1:])
		}
		if _, ok := hash[a]; !ok {
			hash[a] = 0
		}
		hash[a]++
	}
	return max(hash)
}

func Solution(N int, A []int) []int {
	base := 0
	answer := make([]int, N)
	for i := len(A) - 1; i >= 0; i-- {
		idx := A[i]
		if idx == N+1 {
			base = get(N, A[:i])
			break
		}
		answer[idx-1]++
	}
	for idx, _ := range answer {
		answer[idx] += base
	}
	return answer
}
