package LongestIncreasingSubsequence

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func cal(currentNum int, currentSize int, nums, accum []int) int {
	//fmt.Println(currentNum, currentSize, nums, accum)
	if len(nums) == 0 {
		return currentSize
	}
	first := true
	for idx, num := range nums {
		if num > currentNum {
			if !first && (num == accum[len(accum)-1]) {
				break
			}

			first = false

			return max(
				cal(num, currentSize+1, nums[idx+1:], accum[idx+1:]),
				cal(currentNum, currentSize, nums[idx+1:], accum[idx+1:]),
			)
		}
	}
	return currentSize
}

func lengthOfLIS(nums []int) int {
	//fmt.Println(nums)
	accumNums := []int{nums[0]}
	for idx, num := range nums {
		if idx == 0 {
			continue
		}
		if num > accumNums[len(accumNums)-1] {
			accumNums = append(accumNums, num)
		} else {
			accumNums = append(accumNums, accumNums[len(accumNums)-1])
		}
	}
	//fmt.Println(accumNums)
	maxSize := cal(nums[0], 1, nums[1:], accumNums[1:])
	smallerNum := nums[0]

	for idx, num := range nums {
		if smallerNum <= num {
			continue
		} else {
			smallerNum = num
		}
		a := cal(num, 1, nums[idx+1:], accumNums[idx+1:])
		if a > maxSize {
			maxSize = a
		}
	}
	return maxSize
}
