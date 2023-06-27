package MovingAverageFromDataStream

type MovingAverage struct {
	size      int
	total     int
	insertIdx int
	numbers   []int
}

func Constructor(size int) MovingAverage {
	return MovingAverage{
		size: size,
	}
}

func (this *MovingAverage) Next(val int) float64 {

	if this.size > len(this.numbers) {
		this.total += val
		this.numbers = append(this.numbers, val)
	} else {
		lastVal := this.numbers[this.insertIdx]
		this.numbers[this.insertIdx] = val
		this.insertIdx++
		if this.insertIdx == this.size {
			this.insertIdx = 0
		}
		this.total -= lastVal
		this.total += val
	}
	return float64(this.total) / float64(len(this.numbers))
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
