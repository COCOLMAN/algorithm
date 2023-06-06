package ClimbingStairs

var cache map[int]int

func cal(n, current int) int {
	if n == current {
		return 1
	}
	if n < current {
		return 0
	}

	if v, exist := cache[current]; exist {
		return v
	}

	result := cal(n, current+1) + cal(n, current+2)
	cache[current] = result
	return result
}

func climbStairs(n int) int {
	cache = map[int]int{}
	return cal(n, 0)
}
