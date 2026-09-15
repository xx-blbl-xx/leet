package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor236(root, p, q *TreeNode) *TreeNode {
	return findLowestCommonAncestor236(root, p, q)
}

func findLowestCommonAncestor236(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	l := findLowestCommonAncestor236(root.Left, p, q)
	r := findLowestCommonAncestor236(root.Right, p, q)
	if l != nil && r != nil {
		return root
	}

	if l == nil {
		return r
	}

	return l
}
