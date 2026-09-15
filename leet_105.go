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
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	i, j := 0, 0

	for k, v := range inorder {
		if v == preorder[i] {
			j = k
		}
	}

	root.Left = buildTree(preorder[i+1:j+1], inorder[:j])
	root.Right = buildTree(preorder[j+1:], inorder[j+1:])
	return root
}
