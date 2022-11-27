package golang

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	before := head
	up := false
	for l1 != nil && l2 != nil {
		node := &ListNode{}
		val := l1.Val + l2.Val
		if up {
			val++
		}
		if val >= 10 {
			val -= 10
			up = true
		} else {
			up = false
		}
		node.Val = val
		l1 = l1.Next
		l2 = l2.Next

		before.Next = node
		before = node
	}
	for l1 != nil {
		node := &ListNode{}
		val := l1.Val
		if up {
			val++
		}
		if val >= 10 {
			val -= 10
			up = true
		} else {
			up = false
		}
		node.Val = val
		l1 = l1.Next
		before.Next = node
		before = node
	}
	for l2 != nil {
		node := &ListNode{}
		val := l2.Val
		if up {
			val++
		}
		if val >= 10 {
			val -= 10
			up = true
		} else {
			up = false
		}
		node.Val = val
		l2 = l2.Next
		before.Next = node
		before = node
	}
	if up {
		before.Next = &ListNode{1, nil}
	}

	return head.Next
}
