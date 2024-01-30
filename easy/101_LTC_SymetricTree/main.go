package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	leftRoot := root.Left
	rightRoot := root.Right
	for leftRoot != nil || rightRoot != nil {
		if root.Left == nil || root.Right == nil {
			return false
		}
		if leftRoot.Val != rightRoot.Val {
			return false
		}
		leftRoot.
	}
	return true
}
