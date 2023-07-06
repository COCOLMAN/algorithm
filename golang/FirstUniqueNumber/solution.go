package FirstUniqueNumber

type FirstUnique struct {
	uniqueNumbers    []int
	numbersCount     map[int]int
	currentUniqueIdx int
}

func Constructor(nums []int) FirstUnique {
	var uniqueNumbers []int
	numbersCount := map[int]int{}

	for _, number := range nums {
		if _, exist := numbersCount[number]; !exist {
			numbersCount[number] = 0
			uniqueNumbers = append(uniqueNumbers, number)
		}
		numbersCount[number]++
	}

	currentUniqueIdx := -1
	for idx, num := range uniqueNumbers {
		if numbersCount[num] == 1 {
			currentUniqueIdx = idx
			break
		}
	}
	return FirstUnique{
		uniqueNumbers:    uniqueNumbers,
		numbersCount:     numbersCount,
		currentUniqueIdx: currentUniqueIdx,
	}
}

func (this *FirstUnique) ShowFirstUnique() int {
	if this.currentUniqueIdx == -1 {
		return -1
	}
	return this.uniqueNumbers[this.currentUniqueIdx]
}

func (this *FirstUnique) Add(value int) {
	if _, exist := this.numbersCount[value]; !exist {
		this.numbersCount[value] = 0
		this.uniqueNumbers = append(this.uniqueNumbers, value)
	}
	this.numbersCount[value]++

	if this.ShowFirstUnique() != -1 && this.ShowFirstUnique() != value {
		return
	}

	i := 1
	for this.currentUniqueIdx+i < len(this.uniqueNumbers) {
		targetNumber := this.uniqueNumbers[this.currentUniqueIdx+i]
		if this.numbersCount[targetNumber] == 1 {
			this.currentUniqueIdx += i
			return
		}
		i++
	}
	this.currentUniqueIdx = -1
}
