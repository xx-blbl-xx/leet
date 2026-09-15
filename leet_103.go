package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	m := make(map[int][]*TreeNode)
	m[0] = []*TreeNode{root}
	findZigzagLevelOrder(m, 0)

	l := len(m)
	res := [][]int{}
	for i := range l {
		r := []int{}
		if i%2 == 0 {
			for _, v := range m[i] {
				r = append(r, v.Val)
			}
		} else {
			ll := len(m[i]) - 1
			for j := ll; j >= 0; j-- {
				r = append(r, m[i][j].Val)
			}
		}

		res = append(res, r)
	}

	return res
}

func findZigzagLevelOrder(m map[int][]*TreeNode, level int) {
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
		findZigzagLevelOrder(m, level+1)
	}
}
