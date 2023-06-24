package MissingRanges

func findMissingRanges(nums []int, lower int, upper int) [][]int {
	idx := 0
	var result [][]int
	for idx < len(nums) {
		if nums[idx] != lower {
			result = append(result, []int{lower, nums[idx] - 1})
			lower = nums[idx] + 1
			idx++
		} else {
			lower++
			idx++
		}

	}
	if lower <= upper {
		result = append(result, []int{lower, upper})
	}
	return result
}
