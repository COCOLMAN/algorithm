package BuildArrayFromPermutation

import "fmt"

func buildArray(nums []int) []int {
	//ans := make([]int, len(nums))
	//for idx, num := range nums {
	//	ans[idx] = nums[num]
	//}
	//return ans
	//	nums: []int{0, 2, 1, 5, 3, 4},
	/*
		0, 2, 1, 5, 3, 4
		b = qa + r
		r = b // q
		nums[idx] + (q * nums[num])
	*/
	q := len(nums)
	fmt.Println("q is", q)
	/*
		[1, 0]
		1 + (0 % 2) * 2 = 1
		0 + (1 % 2) * 2 = 2
		(1 + (0 % 2) * 2) // 2 = 1
		(0 + (1 % 2) * 2) // 2 = 0
	*/
	for idx, num := range nums {
		//r := num
		//b := nums[num] % q
		//b := nums[num]
		//nums[idx] = q*b + r
		nums[idx] = (nums[num]%q)*q + num
	}
	fmt.Println(nums)
	for idx, _ := range nums {
		nums[idx] /= q
	}
	fmt.Println(nums)
	return nums
}
