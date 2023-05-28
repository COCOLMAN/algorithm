package DeleteAndEarn

import (
	"sort"
)

func deleteAndEarn(nums []int) int {
	scores := map[int]int{}
	for _, num := range nums {
		_, exist := scores[num]
		if !exist {
			scores[num] = 0
		}
		scores[num]++
	}
	max := 0
	points := [][]int{}
	for point, count := range scores {
		points = append(points, []int{point, point * count})
	}

	sort.Slice(points, func(i, j int) bool {
		return points[i][1] > points[j][1]
	})

	globalBan := map[int]bool{}
	for _, candidate := range nums {
		if _, exist := globalBan[candidate]; exist {
			continue
		}
		globalBan[candidate] = true
		result := 0
		ban := map[int]interface{}{}
		if _, exist := scores[candidate-1]; exist {
			ban[candidate-1] = nil
		}
		if _, exist := scores[candidate+1]; exist {
			ban[candidate+1] = nil
		}
		for _, pointInfo := range points {
			scoreKey, totalPoint := pointInfo[0], pointInfo[1]
			if _, exist := ban[scoreKey]; !exist {
				result += totalPoint
				if _, exist := scores[scoreKey-1]; exist {
					ban[scoreKey-1] = nil
				}
				if _, exist := scores[scoreKey+1]; exist {
					ban[scoreKey+1] = nil
				}
			}
		}
		if result > max {
			max = result
		}
	}
	return max
}
