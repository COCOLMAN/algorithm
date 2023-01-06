package SameTree

import "testing"

func Test_isSameTree(t *testing.T) {
	type args struct {
		p *TreeNode
		q *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"default test", args{
			p: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
			q: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
		}, true},
		{"default test", args{
			p: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}},
			q: &TreeNode{Val: 1, Right: &TreeNode{Val: 2}},
		}, false},
		{"default test", args{
			p: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 1}},
			q: &TreeNode{Val: 1, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}},
		}, false},
		{"additional test", args{
			p: &TreeNode{Val: 1, Left: &TreeNode{Val: 1}},
			q: &TreeNode{Val: 1, Left: nil, Right: &TreeNode{Val: 1}},
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSameTree(tt.args.p, tt.args.q); got != tt.want {
				t.Errorf("isSameTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
