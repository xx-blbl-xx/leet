package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// TODO
func maxPathSum(root *TreeNode) int {
	t := root.Val
	l := getMaxPath(root.Left, t)
	r := getMaxPath(root.Right, t)

	return max(t, l, r, l+t, t+r, l+t+r)

}

func getMaxPath(root *TreeNode, res int) int {
	if root == nil {
		return res
	}
	t := root.Val + res
	l := max(getMaxPath(root.Left, t), 0)
	r := max(getMaxPath(root.Right, t), 0)

	return max(t, l, r)
}
