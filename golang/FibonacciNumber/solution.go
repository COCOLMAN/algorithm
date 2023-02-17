package FibonacciNumber

func fibTail(n, b, bb int) int {
	if n == 2 {
		return b + bb
	}
	return fibTail(n-1, b+bb, b)
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fibTail(n, 1, 0)
}
