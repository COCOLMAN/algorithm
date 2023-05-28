package DeleteAndEarn

import (
	"sort"
)

func max(numbers ...int) int {
	m := numbers[0]
	for _, n := range numbers {
		if m < n {
			m = n
		}
	}
	return m
}

var cache map[int]int

func cal(currentIdx int, pointsList *[]int, pointsMap *map[int]int) int {
	if val, exist := cache[currentIdx]; exist {
		return val
	}
	if currentIdx >= len(*pointsList) {
		return 0
	}
	point := (*pointsList)[currentIdx]
	p := (*pointsMap)[point]

	if currentIdx+1 < len(*pointsList) && point+1 == (*pointsList)[currentIdx+1] {
		result := p + max(cal(currentIdx+2, pointsList, pointsMap), cal(currentIdx+3, pointsList, pointsMap))
		cache[currentIdx] = result
		return result
	}
	result := p + max(cal(currentIdx+1, pointsList, pointsMap), cal(currentIdx+2, pointsList, pointsMap))
	cache[currentIdx] = result
	return result
}

func deleteAndEarn(nums []int) int {
	pointCounts := map[int]int{}
	for _, num := range nums {
		_, exist := pointCounts[num]
		if !exist {
			pointCounts[num] = 0
		}
		pointCounts[num]++
	}
	var pointsList []int
	pointsMap := map[int]int{}
	for point, count := range pointCounts {
		pointsList = append(pointsList, point)
		pointsMap[point] = point * count
	}

	sort.Slice(pointsList, func(i, j int) bool {
		return pointsList[i] < pointsList[j]
	})
	cache = map[int]int{}
	a := max(0+cal(0, &pointsList, &pointsMap), 0+cal(1, &pointsList, &pointsMap))
	return a
}
