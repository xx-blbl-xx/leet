package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST98(root *TreeNode) bool {
	r := []int{}
	findisValidBST98(root, &r)

	m := r[0]
	for i := 1; i < len(r); i++ {
		if r[i] > m {
			m = r[i]
			continue
		}
		return false
	}

	return true
}

func findisValidBST98(root *TreeNode, r *[]int) {
	if root == nil {
		return
	}
	l := len(*r)
	if l > 1 && (*r)[l-2] > (*r)[l-1] {
		return
	}

	findisValidBST98(root.Left, r)
	*r = append(*r, root.Val)
	findisValidBST98(root.Right, r)
}
