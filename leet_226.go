package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree226(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	invert226(root)

	return root
}

func invert226(r *TreeNode) {
	if r == nil {
		return
	}

	r.Left, r.Right = r.Right, r.Left

	invert226(r.Left)
	invert226(r.Right)
}
