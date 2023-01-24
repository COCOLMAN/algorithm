package CountDiv

func Solution(A int, B int, K int) int {
	if A > B {
		return 0
	}
	diff := B - A

	if A%K == 0 {
		return 1 + (diff / K)
	}
	return Solution(A+(K-A%K), B, K)
}
