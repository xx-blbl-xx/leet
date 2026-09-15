package main

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
	l := []int{}
	r := 0
	getSumNumbers(root, l, &r)

	return r
}

func getSumNumbers(root *TreeNode, l []int, res *int) {
	if root == nil {
		return
	}
	l = append(l, root.Val)
	if root.Left == nil && root.Right == nil {
		ll := len(l)
		for k, v := range l {
			*res += (v * int(math.Pow10(ll-k-1)))
		}
	} else {
		getSumNumbers(root.Left, l, res)
		getSumNumbers(root.Right, l, res)
	}

}
