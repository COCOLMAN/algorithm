package AddTwoNumbers2

import (
	"reflect"
	"testing"
)

func translateListNode(nums []int) *ListNode {
	head := &ListNode{}
	node := head
	for _, num := range nums {
		current := &ListNode{
			Val: num,
		}
		node.Next = current
		node = current
	}
	return head.Next
}

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 []int
		l2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[]int{7, 2, 4, 3}, []int{5, 6, 4}}, []int{7, 8, 0, 7}},
		{"test1", args{[]int{2, 4, 3}, []int{5, 6, 4}}, []int{8, 0, 7}},
		{"test1", args{[]int{0}, []int{0}}, []int{0}},
		{"test1", args{[]int{1}, []int{9}}, []int{1, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(translateListNode(tt.args.l1), translateListNode(tt.args.l2)); !reflect.DeepEqual(got, translateListNode(tt.want)) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
