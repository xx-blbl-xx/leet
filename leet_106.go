package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree106(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	l := len(postorder)

	root := &TreeNode{Val: postorder[l-1]}

	i := 0
	for ; i < l; i++ {
		if inorder[i] == postorder[l-1] {
			break
		}
	}

	root.Left = buildTree106(inorder[:i], postorder[:i])
	root.Right = buildTree106(inorder[i+1:], postorder[i:l-1])

	return root

}
