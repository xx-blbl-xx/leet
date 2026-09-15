package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth104(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(findMaxDepth(root.Left, 1), findMaxDepth(root.Right, 1))
}

func findMaxDepth(t *TreeNode, i int) int {
	if t == nil {
		return i
	}
	return max(findMaxDepth(t.Left, i+1), findMaxDepth(t.Right, i+1))
}
