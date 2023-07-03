package CheckIfANumberIsMajorityElementInASortedArray

func isMajorityElement(nums []int, target int) bool {
	countMap := map[int]int{}

	for _, num := range nums[:len(nums)+1/2] {
		countMap[num]++
	}

	if count, exist := countMap[target]; exist {
		if count > len(nums)/2 {
			return true
		}
	}
	return false
}
