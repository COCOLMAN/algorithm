package MissingRanges

func findMissingRanges(nums []int, lower int, upper int) [][]int {
	idx := 0
	var result [][]int
	var interval []int
	for idx < len(nums) {
		if nums[idx] != lower {
			if len(interval) == 0 {
				interval = append(interval, lower)
				interval = append(interval, lower)
			} else {
				interval[1] = lower
			}
			lower++
		} else {
			if len(interval) != 0 {
				result = append(result, interval)
				interval = []int{}
			}
			lower++
			idx++
		}
	}
	if lower <= upper {
		result = append(result, []int{lower, upper})
	}
	return result
}
