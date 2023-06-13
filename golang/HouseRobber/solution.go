package HouseRobber

var cache map[int]int

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func _rob(nums []int, currentPosition int, amount int) int {
	if v, exist := cache[currentPosition]; exist {
		return v + amount
	}

	if currentPosition == len(nums)-1 {
		return amount + nums[currentPosition]
	}
	if currentPosition > len(nums)-1 {
		return amount
	}

	result := max(
		_rob(nums, currentPosition+2, amount+nums[currentPosition]),
		_rob(nums, currentPosition+3, amount+nums[currentPosition+1]),
	)
	cache[currentPosition] = result - amount
	return result
}

func rob(nums []int) int {
	cache = map[int]int{}
	return _rob(nums, 0, 0)
}
