package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest230(root *TreeNode, k int) int {
	r := []int{}
	findKthSmallest(root, &r, k)

	return r[k-1]
}

func findKthSmallest(root *TreeNode, r *[]int, k int) {
	if root == nil {
		return
	}
	if len(*r) == k {
		return
	}

	findKthSmallest(root.Left, r, k)
	*r = append(*r, root.Val)
	findKthSmallest(root.Right, r, k)
}
