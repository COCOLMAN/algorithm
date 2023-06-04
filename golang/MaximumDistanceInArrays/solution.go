package MaximumDistanceInArrays

func getSmaller(array []int) int {
	return array[0]
}

func getBigger(array []int) int {
	return array[len(array)-1]
}

func maxDistance(arrays [][]int) int {
	minN1, min1Idx := getSmaller(arrays[0]), 0
	maxN1, max1Idx := getBigger(arrays[0]), 0
	for idx, array := range arrays {
		currentMin := getSmaller(array)
		currentMax := getBigger(array)
		if currentMax > maxN1 {
			maxN1 = currentMax
			max1Idx = idx
		}
		if currentMin < minN1 {
			minN1 = currentMin
			min1Idx = idx
		}
	}

	if min1Idx != max1Idx {
		return maxN1 - minN1
	}
	minN2 := 100000000
	maxN2 := -100000000
	for idx, array := range arrays {
		currentMin := getSmaller(array)
		currentMax := getBigger(array)
		if minN2 >= currentMin && min1Idx != idx {
			minN2 = currentMin

		}
		if maxN2 <= currentMax && max1Idx != idx {
			maxN2 = currentMax

		}
	}

	if maxN1-minN2 > maxN2-minN1 {
		return maxN1 - minN2
	}
	return maxN2 - minN1
}
