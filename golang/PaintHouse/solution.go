package PaintHouse

import "fmt"

var cache map[string]int

var candidate = map[int][]int{
	0: []int{1, 2},
	1: []int{0, 2},
	2: []int{0, 1},
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getKey(a, b int) string {
	return fmt.Sprintf("%d:%d", a, b)
}

func cal(current, ban, idx int, costs [][]int) int {
	key := getKey(ban, idx)
	if val, exist := cache[key]; exist {
		return current + val
	}

	if idx == len(costs) {
		return current
	}

	c := candidate[ban]
	a := c[0]
	b := c[1]
	r1 := cal(current+costs[idx][a], a, idx+1, costs)
	key1 := getKey(a, idx+1)
	cache[key1] = r1 - (current + costs[idx][a])

	r2 := cal(current+costs[idx][b], b, idx+1, costs)
	key2 := getKey(b, idx+1)
	cache[key2] = r2 - (current + costs[idx][b])
	return min(r1, r2)
}

func minCost(costs [][]int) int {
	cache = map[string]int{}
	return min(
		min(
			cal(costs[0][0], 0, 1, costs),
			cal(costs[0][1], 1, 1, costs),
		),
		cal(costs[0][2], 2, 1, costs),
	)
}
