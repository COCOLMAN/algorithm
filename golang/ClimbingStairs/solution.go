package ClimbingStairs

var cache map[int]int

func cal(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	if v, exist := cache[n]; exist {
		return v
	}

	result := cal(n-1) + cal(n-2)
	cache[n] = result
	return result
}

func climbStairs(n int) int {
	cache = map[int]int{}
	return cal(n)
}
