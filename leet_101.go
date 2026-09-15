package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric101(root *TreeNode) bool {
	return checkIsSymmetric101(root, root)
}

func checkIsSymmetric101(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}

	if l == nil || r == nil {
		return false
	}

	if l.Val != r.Val {
		return false
	}

	return checkIsSymmetric101(l.Left, r.Right) && checkIsSymmetric101(l.Right, r.Left)
}
