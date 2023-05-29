package InsertSort

import (
	"fmt"
)

func insertSort(numbers []int) {
	for i := 1; i < len(numbers); i++ {
		num := numbers[i]
		for j := i - 1; j >= 0; j-- {
			if numbers[j] > num {
				tmp := numbers[j]
				numbers[j] = num
				numbers[j+1] = tmp
			} else {
				break
			}
		}
	}
	fmt.Println(numbers)
}
