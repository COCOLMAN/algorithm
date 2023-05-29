package UniquePaths

import "fmt"

type Cache struct {
	data map[string]int
}

var cache Cache

func (c *Cache) set(m, n, val int) {
	c.data[fmt.Sprintf("%d:%d", m, n)] = val
}

func (c *Cache) get(m, n int) (int, bool) {
	val, exist := c.data[fmt.Sprintf("%d:%d", m, n)]
	return val, exist
}

func cal(m int, n int) int {
	if val, exist := cache.get(m, n); exist {
		return val
	}
	if m == 0 || n == 0 {
		return 1
	}
	a := cal(m-1, n)
	cache.set(m-1, n, a)
	b := cal(m, n-1)
	cache.set(m, n-1, b)
	return a + b
}

func uniquePaths(m int, n int) int {
	cache = Cache{
		data: map[string]int{},
	}
	m = m - 1
	n = n - 1
	return cal(m, n)
}
