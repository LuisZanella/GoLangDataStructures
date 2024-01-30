package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
*

	func inorderTraversal(root *TreeNode) []int {
		var result = []int{}

		if root == nil {
			return result
		}

		result = append(result, inorderTraversal(root.Left)...)
		result = append(result, root.Val)
		result = append(result, inorderTraversal(root.Right)...)

		return result
	}
*/
func inorderTraversal(root *TreeNode) []int {
	var result []int
	transverseInOrder(root, &result)
	return result
}

func transverseInOrder(root *TreeNode, result *[]int) {
	if root != nil {
		transverseInOrder(root.Left, result)
		*result = append(*result, root.Val)
		transverseInOrder(root.Right, result)
	}
}

func inorderTraversal2(root *TreeNode) []int {
	var result = []int{}

	if root == nil {
		return result
	}

	result = append(result, inorderTraversal(root.Left)...)
	result = append(result, root.Val)
	result = append(result, inorderTraversal(root.Right)...)

	return result
}
