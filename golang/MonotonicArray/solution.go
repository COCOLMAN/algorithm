package MonotonicArray

func isMonotonic(nums []int) bool {
	if len(nums) <= 2 {
		return true
	}
	firstIdx, secondIdx := 0, 1
	var start, increase bool
	for secondIdx < len(nums) {
		first := nums[firstIdx]
		second := nums[secondIdx]

		if first != second {
			if !start {
				start = true
				if first-second < 0 {
					increase = true
				}
			} else {
				if increase {
					if first > second {
						return false
					}
				} else {
					if second > first {
						return false
					}
				}
			}
		}

		firstIdx++
		secondIdx++
	}

	return true
}
