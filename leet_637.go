package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}
	m := make(map[int][]*TreeNode)
	m[0] = []*TreeNode{root}
	findAverageOfLevels(m, 0)

	l := len(m)
	res := []float64{}
	for i := range l {
		r := 0
		a := 0
		for _, v := range m[i] {
			r += v.Val
			a++
		}
		res = append(res, float64(r)/float64(a))
	}

	return res
}

func findAverageOfLevels(m map[int][]*TreeNode, level int) {
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
		findAverageOfLevels(m, level+1)
	}
}
