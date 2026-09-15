package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView199(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	m := make(map[int][]*TreeNode)
	m[0] = []*TreeNode{root}
	findRightSideView(m, 0)

	l := len(m)
	res := []int{}
	for i := range l {
		ll := len(m[i])
		res = append(res, m[i][ll-1].Val)
	}

	return res
}

func findRightSideView(m map[int][]*TreeNode, level int) {
	res := []*TreeNode{}
	for _, v := range m[level] {
		if v.Left != nil {
			res = append(res, v.Left)
		}
		if v.Right != nil {
			res = append(res, v.Right)
		}
	}

	if len(res) != 0 {
		m[level+1] = res
		findRightSideView(m, level+1)
	}
}
