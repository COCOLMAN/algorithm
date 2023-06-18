package CountingElements

func countElements(arr []int) int {
	countMap := map[int]int{}
	for _, num := range arr {
		if _, exist := countMap[num]; !exist {
			countMap[num] = 0
		}
		countMap[num]++
	}

	result := 0
	for _, num := range arr {
		if _, exist := countMap[num+1]; exist {
			if countMap[num+1] > 0 {
				result++
			}
		}
	}
	return result
}
