package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder102(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	m := make(map[int][]*TreeNode)
	m[0] = []*TreeNode{root}
	findlevelOrder102(m, 0)

	l := len(m)
	res := [][]int{}
	for i := range l {
		r := []int{}

		for _, v := range m[i] {
			r = append(r, v.Val)
		}
		res = append(res, r)
	}

	return res
}

func findlevelOrder102(m map[int][]*TreeNode, level int) {
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
		findlevelOrder102(m, level+1)
	}
}
