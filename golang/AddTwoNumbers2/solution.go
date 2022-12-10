package AddTwoNumbers2

type ListNode struct {
	Val  int
	Next *ListNode
}

func nodeToSlice(l *ListNode) []int {
	numbers := []int{}
	for l != nil {
		numbers = append(numbers, l.Val)
		l = l.Next
	}
	return numbers
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	s1 := nodeToSlice(l1)
	s2 := nodeToSlice(l2)
	p := 0
	up := false
	var last *ListNode
	var current *ListNode
	for p < len(s1) || p < len(s2) {
		var n1, n2 int
		if p < len(s1) {
			n1 = s1[len(s1)-1-p]
		}
		if p < len(s2) {
			n2 = s2[len(s2)-1-p]
		}
		n := n1 + n2
		if up {
			n++
			up = false
		}
		if n >= 10 {
			n -= 10
			up = true
		}
		current = &ListNode{
			Val: n,
		}
		current.Next = last
		last = current
		p++
	}
	if up {
		current = &ListNode{
			Val:  1,
			Next: current,
		}
	}

	return current
}
