package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten114(root *TreeNode) {
	if root == nil {
		return
	}
	r := []*TreeNode{}
	r = preorder(root, r)

	for i := 0; i < len(r)-1; i++ {
		r[i].Left = nil
		r[i].Right = r[i+1]
	}

}

func preorder(root *TreeNode, r []*TreeNode) []*TreeNode {
	if root == nil {
		return r
	}
	r = append(r, root)

	r = preorder(root.Left, r)

	r = preorder(root.Right, r)

	return r
}
