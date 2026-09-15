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
func getMinimumDifference(root *TreeNode) int {
	res := []int{}
	getMinimumDiff(root.Left, &res)
	res = append(res, root.Val)
	getMinimumDiff(root.Right, &res)

	r := math.MaxInt
	for i := 1; i < len(res); i++ {
		r = min(r, res[i]-res[i-1])
	}

	return r
}

func getMinimumDiff(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	getMinimumDiff(root.Left, res)
	*res = append(*res, root.Val)
	getMinimumDiff(root.Right, res)
}
