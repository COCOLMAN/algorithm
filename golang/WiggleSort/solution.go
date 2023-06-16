package WiggleSort

func swap(i, j int, nums *[]int) {
	tmp := (*nums)[i]
	(*nums)[i] = (*nums)[j]
	(*nums)[j] = tmp
}

func wiggleSort(nums []int) {
	if len(nums) == 1 {
		return
	}
	i := 0
	j := 1
	for j < len(nums) {
		if i%2 == 0 {
			// even; next number must big
			if nums[i] > nums[j] {
				swap(i, j, &nums)
			}
		} else {
			// odd; next number must small
			if nums[i] < nums[j] {
				swap(i, j, &nums)
			}
		}
		i++
		j++
	}
}
