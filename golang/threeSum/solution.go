package threeSum

import (
	"fmt"
	"sort"
)

func contains(a int, s []int) bool {
	for _, r := range s {
		if r == a {
			return true
		}
	}
	return false
}

func twoSum(goal int, nums []int, indexMap map[int][]int) [][]int {
	var answer [][]int
	for idx, num := range nums {
		target := goal - num
		indexList, exist := indexMap[target]
		if !exist {
			continue
		}

		for _, idxItem := range indexList {
			if idx == idxItem || idx > idxItem {
				continue
			}
			answer = append(answer, []int{idx, idxItem})
		}
	}
	return answer
}

func makeKey(nums ...int) string {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	return fmt.Sprintf("%d:%d:%d", nums[0], nums[1], nums[2])
}

func threeSum(nums []int) [][]int {
	indexMap := map[int][]int{}
	for idx, num := range nums {
		if _, exist := indexMap[num]; !exist {
			indexMap[num] = []int{}
		}
		indexMap[num] = append(indexMap[num], idx)
	}

	var answer [][]int
	existCheck := map[string]bool{}
	for idx, num := range nums {
		want := 0 - num
		result := twoSum(want, nums, indexMap)
		for _, item := range result {
			if !contains(idx, item) && item[1] < idx {
				newItem := append(item, idx)
				key := makeKey(nums[newItem[0]], nums[newItem[1]], nums[newItem[2]])
				_, exist := existCheck[key]
				if exist {
					continue
				}

				existCheck[key] = true
				answer = append(answer, []int{nums[newItem[0]], nums[newItem[1]], nums[newItem[2]]})
			}
		}
	}
	return answer
}
