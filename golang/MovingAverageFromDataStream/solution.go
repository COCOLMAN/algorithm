package MovingAverageFromDataStream

type MovingAverage struct {
	size    int
	total   int
	numbers []int
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
		lastVal := this.numbers[0]
		this.numbers = this.numbers[1:]
		this.numbers = append(this.numbers, val)
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
