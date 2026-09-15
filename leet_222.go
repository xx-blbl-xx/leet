package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
	return countN(root)
}

func countN(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := countN(root.Left)
	r := countN(root.Right)

	return l + r + 1

}
