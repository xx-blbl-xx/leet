package main

import "fmt"

func numEquivDominoPairs(dominoes [][]int) int {
	mm := make(map[string]int, 0)

	for _, d := range dominoes {
		ds := ""
		if d[0] > d[1] {
			ds = fmt.Sprintf("%d%d", d[0], d[1])
		} else {
			ds = fmt.Sprintf("%d%d", d[1], d[0])
		}
		mm[ds]++
	}
	res := 0

	for _, v := range mm {
		if v == 1 {
			continue
		}
		res += v * (v - 1) / 2
	}

	return res
}
