package main

/**
 * Definition for a Node.
 * type Node117 struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */
type Node117 struct {
	Val   int
	Left  *Node117
	Right *Node117
	Next  *Node117
}

func connect(root *Node117) *Node117 {
	if root == nil {
		return nil
	}
	mm := [][]*Node117{{root}}
	con2(1, &mm)

	for _, m := range mm {
		for k, v := range m {
			if k == 0 {
				continue
			}
			m[k-1].Next = v
		}
	}

	return root
}

func con2(depth int, mm *[][]*Node117) {
	if len((*mm)[depth-1]) == 0 {
		return
	}
	(*mm) = append((*mm), make([]*Node117, 0))
	for _, r := range (*mm)[depth-1] {
		if r.Left != nil {
			(*mm)[depth] = append((*mm)[depth], r.Left)
		}
		if r.Right != nil {
			(*mm)[depth] = append((*mm)[depth], r.Right)
		}
	}

	con2(depth+1, mm)
}
