package RemoveElement

type Queue struct {
	items  []int
	cursor int
}

func (q *Queue) Append(item int) {
	q.items = append(q.items, item)
}

func (q *Queue) Pop() int {
	if q.cursor >= len(q.items) {
		return -1
	}
	item := q.items[q.cursor]
	q.cursor++
	return item
}

func removeElement(nums []int, val int) int {
	queue := Queue{}
	count := 0
	for idx, num := range nums {
		if num == val {
			queue.Append(idx)
			continue
		}

		var newIdx int
		queueingIdx := queue.Pop()
		if queueingIdx == -1 {
			newIdx = idx
		} else {
			newIdx = queueingIdx
			queue.Append(idx)
		}
		nums[newIdx] = num
		count++
	}
	return count
}
