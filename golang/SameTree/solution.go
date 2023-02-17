package SameTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getPreorder(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return getPreorder(left.Left, right.Left) && getPreorder(left.Right, right.Right)
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	return getPreorder(p, q)
}
