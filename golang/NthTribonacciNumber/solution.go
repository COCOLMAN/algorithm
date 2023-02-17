package NthTribonacciNumber

func triTail(n, b, bb, bbb int) int {
	current := b + bb + bbb
	if n == 3 {
		return current
	}
	return triTail(n-1, current, b, bb)
}

func tribonacci(n int) int {
	if n < 2 {
		return n
	}
	if n == 2 {
		return 1
	}
	return triTail(n, 1, 1, 0)
}
