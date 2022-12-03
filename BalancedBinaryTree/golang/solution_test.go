package golang

import (
	"testing"
)

func Test_isBalanced(t *testing.T) {
	type args struct {
		root *TreeNode
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test 1", args{root: &TreeNode{3, &TreeNode{Val: 9}, &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}}, true},
		{"test 2", args{root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 2}}}, false},
		{"test 3", args{root: nil}, true},
		{"test 4", args{root: &TreeNode{Val: 1}}, true},
		{"test 5", args{root: &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}}}, false},
		{"test 6", args{root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 4}}}, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4}}}}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBalanced(tt.args.root); got != tt.want {
				t.Errorf("isBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
