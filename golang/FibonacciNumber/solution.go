package FibonacciNumber

func fib(n int) int {
	if n < 2 {
		return n
	}
	answer := 0
	current := 1
	before := 1
	bbefore := 0
	for n != current {
		answer = before + bbefore
		bbefore = before
		before = answer
		current++
	}
	return answer

}
