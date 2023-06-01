package MaximumDistanceInArrays

import "fmt"

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func maxDistance(arrays [][]int) int {
	var diff []int
	var diffResult int
	if abs(arrays[0][0]-arrays[1][len(arrays[1])-1]) > abs(arrays[1][0]-arrays[0][len(arrays[0])-1]) {
		diff = []int{arrays[0][0], arrays[1][len(arrays[1])-1]}
		diffResult = abs(arrays[0][0] - arrays[1][len(arrays[1])-1])
	} else {
		diff = []int{arrays[1][0], arrays[0][len(arrays[0])-1]}
		diffResult = abs(arrays[1][0] - arrays[0][len(arrays[0])-1])
	}
	for _, numbers := range arrays[2:] {
		min := numbers[0]
		max := numbers[len(numbers)-1]
		var newMin, newMax bool
		if abs(min-diff[1]) > diffResult {
			newMin = true
			fmt.Println("newMin", min, diff[1])
		}
		if abs(max-diff[0]) > diffResult {
			newMin = true
			fmt.Println("newMax")
		}
		if newMax && newMin {
			if abs(max-diff[0]) > abs(min-diff[1]) {
				diff[1] = max
				diffResult = abs(max - diff[0])
			} else {
				diff[0] = min
				diffResult = abs(min - diff[1])
			}
		} else if newMax {
			diff[1] = max
			diffResult = abs(max - diff[0])
		} else if newMin {
			diff[0] = min
			diffResult = abs(min - diff[1])
		}
	}

	return diffResult
}
